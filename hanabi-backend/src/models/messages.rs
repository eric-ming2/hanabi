use tokio::sync::mpsc::{Receiver, Sender};

#[derive(Debug)]
pub enum TaskMessage {
    // Messages that a client thread sends to the main thread
    InitClient(String, Sender<(String, TaskMessage)>),
    CloseClient(String),
    StartGame,
    // Messages that the main thread sends to a client thread
    UpdateGameState(),
}
