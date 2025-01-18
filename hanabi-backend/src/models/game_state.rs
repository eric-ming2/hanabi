pub struct GameState {
    players: Vec<PlayerCards>,
    turn: u8,
    deck: Vec<Card>,
    discard_pile: Vec<Card>,
    hints: u8,
    bombs: u8,
    fireworks: HashMap<CardColor, u8>
}

pub struct PlayerCards {
    cards: Vec<(Card, UnknownCard)>
}

