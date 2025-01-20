mod client;
mod models;

use models::game_state::{GameState, GameStatePerspective};
use models::messages::TaskMessage;

use std::collections::HashMap;
use std::sync::Arc;

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
    tx: Sender<(String, TaskMessage)>,
}

#[derive(Debug)]
struct ServerState {
    state: State,
    clients: HashMap<String, ClientInfo>,
    game_state: Option<GameState>,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "127.0.0.1:8080";
    let listener = TcpListener::bind(addr).await?;
    println!("Listening on: {}", addr);

    let mut server_state = ServerState {
        state: State::Lobby,
        clients: HashMap::new(),
        game_state: None,
    };

    let (main_tx, mut main_rx) = mpsc::channel::<(String, TaskMessage)>(100);

    tokio::spawn(async move {
        while let Some((id, msg)) = main_rx.recv().await {
            println!("Main thread received from {}: {:?}", id, msg);
            match msg {
                TaskMessage::InitClient(id, username, tx) => {
                    server_state
                        .clients
                        .insert(id, ClientInfo {
                            player_index: server_state.clients.len() as u8,
                            username,
                            tx,
                        });
                    println!("size: {}", server_state.clients.len());
                }
                TaskMessage::CloseClient(id) => {
                    server_state.clients.remove(&id);
                }
                TaskMessage::StartGame => {
                    if server_state.clients.len() > 2 {
                        server_state.state = State::Game;
                        server_state.game_state =
                            Some(GameState::new(server_state.clients.len() as u8));
                        for (_, client_info) in &server_state.clients {
                            client_info.tx.send((
                                String::from("main"),
                                TaskMessage::UpdateGameState(GameStatePerspective::from_state(
                                    server_state.game_state.clone().unwrap(),
                                    client_info.player_index,
                                )),
                            ))
                            .await
                            .unwrap();
                        }
                    } else {
                        println!(
                            "Need 3 players to start the game. Tried to start with {}",
                            server_state.clients.len()
                        );
                    }
                }
                TaskMessage::UpdateGameState(_) => {
                    unreachable!();
                }
            }
        }
    });

    loop {
        let (stream, _) = listener.accept().await?;
        println!("New connection established");
        tokio::spawn(client::handle_connection(stream, main_tx.clone()));
    }
}
