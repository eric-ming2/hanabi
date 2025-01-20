use super::card::{Card, CardColor, UnknownCard};
use rand::seq::SliceRandom;
use rand::thread_rng;
use std::collections::HashMap;

mod generated {
    pub mod responses {
        include!(concat!(
            env!("CARGO_MANIFEST_DIR"),
            "/src/generated/responses.rs"
        ));
    }
}

use generated::responses::{
    Card as ProtoCard, CardColor as ProtoCardColor, PlayerCards as ProtoPlayerCards, Response,
    UnknownCard as ProtoUnknownCard, UpdateGameResponse,
};

#[derive(Debug, Clone)]
pub struct GameState {
    players: Vec<PlayerCards>,
    turn: u8,
    deck: Vec<Card>,
    discard_pile: Vec<Card>,
    hints: u8,
    bombs: u8,
    fireworks: HashMap<CardColor, u8>,
}

#[derive(Debug, Clone)]
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

#[derive(Debug)]
pub struct GameStatePerspective {
    my_hand: [UnknownCard; 4],
    other_hands: Vec<PlayerCards>,
    turn: u8,
    deck: Vec<Card>,
    discard_pile: Vec<Card>,
    hints: u8,
    bombs: u8,
    fireworks: HashMap<CardColor, u8>,
}

impl GameStatePerspective {
    pub fn from_state(game_state: GameState, player_index: u8) -> Self {
        let my_cards = game_state
            .players
            .get(player_index as usize)
            .unwrap()
            .cards
            .clone();
        let my_hand = [
            my_cards[0].1.clone(),
            my_cards[1].1.clone(),
            my_cards[2].1.clone(),
            my_cards[3].1.clone(),
        ];
        let mut other_hands = game_state.players.clone();
        other_hands.remove(player_index as usize);
        other_hands.rotate_left(player_index as usize);

        GameStatePerspective {
            my_hand,
            other_hands,
            turn: game_state.turn,
            deck: game_state.deck,
            discard_pile: game_state.discard_pile,
            hints: game_state.hints,
            bombs: game_state.bombs,
            fireworks: game_state.fireworks,
        }
    }

    pub fn to_proto(&self) -> Response {
        let update_game_req = UpdateGameResponse {
            my_hand: self.my_hand.iter().map(|c| c.clone().into()).collect(),
            other_hands: self.other_hands.iter().map(|p| p.clone().into()).collect(),
            turn: self.turn as i32,
            deck: self.deck.iter().map(|c| c.clone().into()).collect(),
            discard_pile: self.discard_pile.iter().map(|c| c.clone().into()).collect(),
            hints: self.hints as i32,
            bombs: self.bombs as i32,
            fireworks: self
                .fireworks
                .iter()
                .map(|(k, v)| (k.clone() as i32, *v as i32))
                .collect(),
        };
        Response {
            response_type: 1,
            response: Some(generated::responses::response::Response::UpdateGame(
                update_game_req,
            )),
        }
    }
}

impl From<PlayerCards> for ProtoPlayerCards {
    fn from(pc: PlayerCards) -> Self {
        ProtoPlayerCards {
            cards: pc.cards.iter().map(|(c, _)| c.clone().into()).collect(),
            unknown_cards: pc.cards.iter().map(|(_, uc)| uc.clone().into()).collect(),
        }
    }
}

impl From<Card> for ProtoCard {
    fn from(c: Card) -> Self {
        ProtoCard {
            num: c.num as i32,
            color: match c.color {
                CardColor::White => ProtoCardColor::White.into(),
                CardColor::Yellow => ProtoCardColor::Yellow.into(),
                CardColor::Green => ProtoCardColor::Green.into(),
                CardColor::Blue => ProtoCardColor::Blue.into(),
                CardColor::Red => ProtoCardColor::Red.into(),
            },
        }
    }
}

impl From<UnknownCard> for ProtoUnknownCard {
    fn from(uc: UnknownCard) -> Self {
        ProtoUnknownCard {
            num: uc.num.map(|num| num as i32),
            color: uc.color.map(|color| match color {
                CardColor::White => ProtoCardColor::White.into(),
                CardColor::Yellow => ProtoCardColor::Yellow.into(),
                CardColor::Green => ProtoCardColor::Green.into(),
                CardColor::Blue => ProtoCardColor::Blue.into(),
                CardColor::Red => ProtoCardColor::Red.into(),
            }),
        }
    }
}
