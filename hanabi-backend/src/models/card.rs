#[derive(PartialEq, Eq, Hash, Debug)]
pub enum CardColor {
    White,
    Yellow,
    Green,
    Blue,
    Red,
}

#[derive(Debug)]
pub struct Card {
    pub num: u8, // 1-5
    pub color: CardColor,
}

#[derive(Debug)]
pub struct UnknownCard {
    pub num: Option<u8>,
    pub color: Option<CardColor>,
}
