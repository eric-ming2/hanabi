use tokio::sync::mpsc::{Receiver, Sender};

use super::game_state::{GamePerspective, StartedGamePerspective};

#[derive(Debug)]
pub enum TaskMessage {
    // Messages that a client thread sends to the main thread
    InitClient(String, String, Sender<TaskMessage>), // id, username, tx. Could refactor to struct?
    CloseClient(String),
    Ready,
    StartGame,
    // Messages that the main thread sends to a client thread
    UpdateGamePerspective(GamePerspective),
}
