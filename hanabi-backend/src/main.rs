mod client;
mod models;

use models::game_state::{GameState, GameStatePerspective};
use models::messages::TaskMessage;
use tokio_tungstenite::tungstenite::handshake::server;

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
struct ServerState {
    state: State,
    clients: HashMap<String, (u8, Sender<(String, TaskMessage)>)>,
    game_state: Option<GameState>,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Bind the TCP listener to an address and start listening for connections
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
                TaskMessage::InitClient(id, tx) => {
                    server_state
                        .clients
                        .insert(id, (server_state.clients.len() as u8, tx));
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
                        for (_, (player_index, tx)) in &server_state.clients {
                            tx.send((
                                String::from("main"),
                                TaskMessage::UpdateGameState(GameStatePerspective::from_state(
                                    // TODO: Can I pass a reference instead of cloning here?
                                    server_state.game_state.clone().unwrap(),
                                    *player_index,
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
        // Accept incoming TCP connections
        let (stream, _) = listener.accept().await?;
        println!("New connection established");
        // Spawn a new task to handle the WebSocket connection
        tokio::spawn(client::handle_connection(stream, main_tx.clone()));
    }
}
