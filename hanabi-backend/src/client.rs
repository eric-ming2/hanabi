use std::sync::Arc;

use futures::{sink::SinkExt, stream::StreamExt};
use mpsc::{channel, Sender};
use tokio::sync::{mpsc, Mutex};
use tokio_tungstenite::tungstenite::protocol::Message;
use tokio_tungstenite::{accept_async, tungstenite::Error};
use uuid::Uuid;

use crate::models::messages::TaskMessage;

// Function to handle the WebSocket connection
pub(crate) async fn handle_connection(
    stream: tokio::net::TcpStream,
    main_tx: Sender<(String, TaskMessage)>,
) {
    if let Err(e) = process_connection(stream, main_tx).await {
        eprintln!("Error during connection: {}", e);
    }
}

async fn process_connection(
    stream: tokio::net::TcpStream,
    main_tx: Sender<(String, TaskMessage)>,
) -> Result<(), Error> {
    let ws_stream = accept_async(stream)
        .await
        .expect("Failed to complete WebSocket handshake");

    println!("WebSocket connection estabished");

    // Split the WebSocket stream into a sender and receiver
    let (mut write, mut read) = ws_stream.split();

    let (tx, mut rx) = channel::<(String, TaskMessage)>(100);

    tokio::spawn(async move {
        while let Some((_, msg)) = rx.recv().await {
            match msg {
                TaskMessage::UpdateGameState(game_state_perspective) => {
                    write
                        .send(Message::Text(
                            format!("{:?}", game_state_perspective).into(),
                        ))
                        .await
                        .unwrap();
                }
                _ => {
                    unreachable!()
                }
            }
        }
    });

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
                let msg_string = msg.as_str();
                match msg_string {
                    "StartGame" => {
                        main_tx
                            .send((id.clone(), TaskMessage::StartGame))
                            .await
                            .unwrap();
                    }
                    _ => {
                        println!("Invalid message received");
                    }
                }
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
