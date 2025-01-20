use futures::{sink::SinkExt, stream::StreamExt};
use mpsc::{channel, Sender};
use prost::Message as ProstMessage;
use tokio::sync::{mpsc, Mutex};

use tokio_tungstenite::tungstenite::protocol::Message;
use tokio_tungstenite::tungstenite::Bytes;
use tokio_tungstenite::{accept_async, tungstenite::Error};

use crate::models::messages::TaskMessage;

mod generated {
    pub mod requests {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/src/generated/requests.rs"
        ));
    }
}

use crate::client::generated::requests::request::Request::InitConnection;
use crate::client::generated::requests::{Request, RequestType};

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

    // TODO: This code smells
    let mut id = "uninitialized".to_string();

    tokio::spawn(async move {
        while let Some((_, msg)) = rx.recv().await {
            match msg {
                TaskMessage::UpdateGameState(game_state_perspective) => {
                    write
                        .send(Message::Binary(Bytes::from(
                            game_state_perspective.to_proto().encode_to_vec(),
                        )))
                        .await
                        .unwrap();
                }
                _ => {
                    unreachable!()
                }
            }
        }
    });

    // Process messages from the client
    while let Some(message) = read.next().await {
        match message {
            Ok(Message::Binary(msg_bytes)) => {
                // TODO: Error handle better here
                let request = Request::decode(msg_bytes.as_ref()).unwrap();
                match RequestType::try_from(request.request_type).unwrap() {
                    RequestType::InitConnection => {
                        println!("Received InitConnection request");
                        match request.request {
                            Some(InitConnection(init_connection)) => {
                                id = init_connection.id.clone();
                                main_tx
                                    .send((
                                        id.clone(),
                                        TaskMessage::InitClient(id.clone(), init_connection.username.clone(), tx.clone()),
                                    ))
                                    .await
                                    .unwrap();
                            }
                            _ => {
                                unreachable!();
                            }
                        }
                    }
                    RequestType::StartGame => {
                        println!("Received StartGame request");
                        main_tx
                            .send((id.clone(), TaskMessage::StartGame))
                            .await
                            .unwrap();
                    }
                    RequestType::DiscardCard => {
                        println!("Received DiscardCard request");
                    }
                    RequestType::PlayCard => {
                        println!("Received PlayCard request");
                    }
                    RequestType::GiveHint => {
                        println!("Received GiveHint request");
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
