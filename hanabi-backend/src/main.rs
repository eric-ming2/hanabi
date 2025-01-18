use std::collections::HashMap;
use std::sync::Arc;

use futures::{stream::StreamExt, sink::SinkExt};
use tokio::net::TcpListener;
use tokio::sync::{mpsc, Mutex};
use tokio_tungstenite::tungstenite::protocol::Message;
use tokio_tungstenite::{accept_async, tungstenite::Error};
use mpsc::{ Sender, Receiver };

enum State {
    Lobby,
    Game,
}

struct ServerState {
    state: State,
    clients: HashMap<String, mpsc::Sender<(String, String)>>,
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    // Bind the TCP listener to an address and start listening for connections
    let addr = "127.0.0.1:8080";
    let listener = TcpListener::bind(addr).await?;
    println!("Listening on: {}", addr);

    let server_state = ServerState {
        state: State::Lobby,
        clients: HashMap::new(),
    };

    let (main_tx, mut main_rx) = mpsc::channel::<(String, String)>(100);

    tokio::spawn(async move {
        while let Some((id, msg)) = main_rx.recv().await {
            println!("Main thread received from {}: {}", id, msg);
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
async fn handle_connection(stream: tokio::net::TcpStream, main_tx: Sender<(String, String)>) {
    if let Err(e) = process_connection(stream, main_tx).await {
        eprintln!("Error during connection: {}", e);
    }
}

// This function upgrades the TCP stream to a WebSocket and handles the messages
async fn process_connection(stream: tokio::net::TcpStream, main_tx: Sender<(String, String)>) -> Result<(), Error> {
    // Perform the WebSocket handshake (upgrade the TCP stream to a WebSocket)
    let ws_stream = accept_async(stream)
        .await
        .expect("Failed to complete WebSocket handshake");

    println!("WebSocket connection established");

    // Split the WebSocket stream into a sender and receiver
    let (mut write, mut read) = ws_stream.split();

    let (tx, mut rx) = mpsc::channel::<(String, String)>(100);
    let id = "Hi I'm eric".to_string();
    main_tx.send((id.clone(), "initialized eric".to_string())).await.unwrap();

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
