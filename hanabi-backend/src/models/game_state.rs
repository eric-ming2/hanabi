use super::card::{Card, CardColor, UnknownCard};
use rand::seq::SliceRandom;
use rand::thread_rng;
use std::collections::HashMap;

#[derive(Debug)]
pub struct GameState {
    players: Vec<PlayerCards>,
    turn: u8,
    deck: Vec<Card>,
    discard_pile: Vec<Card>,
    hints: u8,
    bombs: u8,
    fireworks: HashMap<CardColor, u8>,
}

#[derive(Debug)]
pub struct PlayerCards {
    cards: [(Card, UnknownCard); 4],
}

impl GameState {
    pub fn new(num_players: u8) -> Self {
        let deck = Self::new_deck();
        let (deck, players) = Self::deal(deck, num_players);
        GameState {
            players,
            turn: 0,
            deck,
            discard_pile: Vec::new(),
            hints: 8,
            bombs: 3,
            fireworks: Self::new_fireworks(),
        }
    }

    fn deal(mut deck: Vec<Card>, num_players: u8) -> (Vec<Card>, Vec<PlayerCards>) {
        let mut players = Vec::new();
        for _ in 0..num_players {
            let cards: [(Card, UnknownCard); 4] = [
                (
                    deck.pop().unwrap(),
                    UnknownCard {
                        num: None,
                        color: None,
                    },
                ),
                (
                    deck.pop().unwrap(),
                    UnknownCard {
                        num: None,
                        color: None,
                    },
                ),
                (
                    deck.pop().unwrap(),
                    UnknownCard {
                        num: None,
                        color: None,
                    },
                ),
                (
                    deck.pop().unwrap(),
                    UnknownCard {
                        num: None,
                        color: None,
                    },
                ),
            ];
            players.push(PlayerCards { cards });
        }
        (deck, players)
    }

    fn new_fireworks() -> HashMap<CardColor, u8> {
        let mut fireworks = HashMap::new();
        fireworks.insert(CardColor::White, 0);
        fireworks.insert(CardColor::Yellow, 0);
        fireworks.insert(CardColor::Green, 0);
        fireworks.insert(CardColor::Blue, 0);
        fireworks.insert(CardColor::Red, 0);
        fireworks
    }

    fn new_deck() -> Vec<Card> {
        let mut deck = Vec::new();
        for _ in 0..3 {
            deck.push(Card {
                num: 1,
                color: CardColor::White,
            });
            deck.push(Card {
                num: 1,
                color: CardColor::Yellow,
            });
            deck.push(Card {
                num: 1,
                color: CardColor::Green,
            });
            deck.push(Card {
                num: 1,
                color: CardColor::Blue,
            });
            deck.push(Card {
                num: 1,
                color: CardColor::Red,
            });
        }
        for _ in 0..2 {
            for i in 2..=4 {
                deck.push(Card {
                    num: i,
                    color: CardColor::White,
                });
                deck.push(Card {
                    num: i,
                    color: CardColor::Yellow,
                });
                deck.push(Card {
                    num: i,
                    color: CardColor::Green,
                });
                deck.push(Card {
                    num: i,
                    color: CardColor::Blue,
                });
                deck.push(Card {
                    num: i,
                    color: CardColor::Red,
                });
            }
        }
        deck.push(Card {
            num: 5,
            color: CardColor::White,
        });
        deck.push(Card {
            num: 5,
            color: CardColor::Yellow,
        });
        deck.push(Card {
            num: 5,
            color: CardColor::Green,
        });
        deck.push(Card {
            num: 5,
            color: CardColor::Blue,
        });
        deck.push(Card {
            num: 5,
            color: CardColor::Red,
        });
        let mut rng = thread_rng();
        deck.shuffle(&mut rng);
        deck
    }
}
