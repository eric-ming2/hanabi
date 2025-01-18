mod models;

use models::game_state::GameState;
use models::messages::TaskMessage;

use std::collections::HashMap;
use std::sync::Arc;

use futures::{sink::SinkExt, stream::StreamExt};
use mpsc::{Receiver, Sender};
use tokio::net::TcpListener;
use tokio::sync::{mpsc, Mutex};
use tokio_tungstenite::tungstenite::protocol::Message;
use tokio_tungstenite::{accept_async, tungstenite::Error};
use uuid::Uuid;

#[derive(Debug)]
enum State {
    Lobby,
    Game,
}

#[derive(Debug)]
struct ServerState {
    state: State,
    clients: HashMap<String, mpsc::Sender<(String, TaskMessage)>>,
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
                    server_state.clients.insert(id, tx);
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
                        println!("Server state: {:?}", server_state)
                    } else {
                        println!(
                            "Need 3 players to start the game. Tried to start with {}",
                            server_state.clients.len()
                        );
                    }
                }
                TaskMessage::UpdateGameState() => todo!(),
            }
        }
    });

    loop {
        // Accept incoming TCP connections
        let (stream, _) = listener.accept().await?;
        println!("New connection established");
        // Spawn a new task to handle the WebSocket connection
        tokio::spawn(handle_connection(stream, main_tx.clone()));
    }
}

// Function to handle the WebSocket connection
async fn handle_connection(stream: tokio::net::TcpStream, main_tx: Sender<(String, TaskMessage)>) {
    if let Err(e) = process_connection(stream, main_tx).await {
        eprintln!("Error during connection: {}", e);
    }
}

// This function upgrades the TCP stream to a WebSocket and handles the messages
async fn process_connection(
    stream: tokio::net::TcpStream,
    main_tx: Sender<(String, TaskMessage)>,
) -> Result<(), Error> {
    // Perform the WebSocket handshake (upgrade the TCP stream to a WebSocket)
    let ws_stream = accept_async(stream)
        .await
        .expect("Failed to complete WebSocket handshake");

    println!("WebSocket connection established");

    // Split the WebSocket stream into a sender and receiver
    let (mut write, mut read) = ws_stream.split();

    let (tx, mut rx) = mpsc::channel::<(String, TaskMessage)>(100);
    let id = Uuid::new_v4().to_string();
    main_tx
        .send((id.clone(), TaskMessage::InitClient(id.clone(), tx)))
        .await
        .unwrap();

    // Process messages from the client
    while let Some(message) = read.next().await {
        match message {
            Ok(Message::Text(msg)) => {
                println!("Received message: {}", msg);
                // Echo the message back to the client
                write.send(Message::Text(msg)).await?;
            }
            Ok(Message::Ping(ping)) => {
                // Handle Ping messages (can respond with Pong)
                write.send(Message::Pong(ping)).await?;
            }
            Ok(Message::Close(close_frame)) => {
                main_tx
                    .send((id.clone(), TaskMessage::CloseClient(id.clone())))
                    .await
                    .unwrap();
                println!("Closing connection: {:?}", close_frame);
                break;
            }
            Err(e) => {
                eprintln!("Error while processing message: {:?}", e);
                break;
            }
            _ => {}
        }
    }

    Ok(())
}
