mod client;
mod models;

use models::game_state::{StartedGamePerspective, StartedGameState};
use models::messages::TaskMessage;

use std::collections::HashMap;
use std::sync::Arc;

use crate::models::game_state::{
    Game, GamePerspective, GameState, NotStartedGameState, Perspective,
};
use mpsc::{channel, Sender};
use tokio::net::TcpListener;
use tokio::sync::{mpsc, Mutex};

#[derive(Debug)]
enum State {
    Lobby,
    Game,
}

#[derive(Debug)]
struct ClientInfo {
    player_index: u8,
    username: String,
    tx: Sender<TaskMessage>,
}

#[derive(Debug)]
struct ServerState {
    clients: HashMap<String, ClientInfo>,
    game: Game,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "127.0.0.1:8080";
    let listener = TcpListener::bind(addr).await?;
    println!("Listening on: {}", addr);

    let mut server_state = ServerState {
        clients: HashMap::new(),
        game: Game::new(),
    };

    let (main_tx, mut main_rx) = mpsc::channel::<(String, TaskMessage)>(100);

    tokio::spawn(async move {
        while let Some((id, msg)) = main_rx.recv().await {
            println!("Main thread received from {}: {:?}", id, msg);
            match msg {
                TaskMessage::InitClient(id, username, tx) => {
                    server_state.game.add_player(username.clone(), id.clone());
                    server_state.clients.insert(
                        id,
                        ClientInfo {
                            player_index: server_state.clients.len() as u8,
                            username,
                            tx,
                        },
                    );
                }
                TaskMessage::CloseClient(id) => {
                    server_state.clients.remove(&id);
                }
                TaskMessage::Ready => {
                    server_state.game.ready(id);
                }
                TaskMessage::StartGame => {
                    // TODO: Should probably move logic inside the Game struct.
                    match &server_state.game.game_state {
                        GameState::NotStarted(nsgs) => {
                            if nsgs.players.len() > 2 && nsgs.players.iter().all(|p| p.ready) {
                                server_state.game.start_game();
                                for client in server_state.clients.values() {
                                    client
                                        .tx
                                        .send(TaskMessage::UpdateGamePerspective(
                                            GamePerspective::from_state(
                                                &server_state.game,
                                                client.player_index,
                                            ),
                                        ))
                                        .await
                                        .expect("Something is wrong");
                                }
                            } else {
                                // TODO: You should send a message to the UI somehow
                                println!("Failed to start game");
                            }
                        }
                        GameState::Started(_) => {
                            println!("Tried to start an already started game.");
                        }
                    }
                }
                TaskMessage::UpdateGamePerspective(_) => {
                    unreachable!();
                }
            }
            for client in server_state.clients.values() {
                client
                    .tx
                    .send(TaskMessage::UpdateGamePerspective(
                        GamePerspective::from_state(&server_state.game, client.player_index),
                    ))
                    .await
                    .expect("Something is wrong");
            }
            println!("game state: {:?}", server_state.game);
        }
    });

    loop {
        let (stream, _) = listener.accept().await?;
        println!("New connection established");
        tokio::spawn(client::handle_connection(stream, main_tx.clone()));
    }
}
