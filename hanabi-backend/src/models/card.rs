pub enum CardColor {
    White,
    Yellow,
    Green,
    Blue,
    Red,
}

pub struct Card {
    pub num: u8, // 1-5
    pub color: CardColor,
}

pub struct UnknownCard {
    pub num: Option<u8>,
    pub color: Option<CardColor>,
}