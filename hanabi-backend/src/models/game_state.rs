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
    update_game_response::GamePerspective as ProtoGamePerspective, Card as ProtoCard,
    CardColor as ProtoCardColor, NotStartedGamePerspective as ProtoNotStartedGamePerspective,
    NotStartedPlayer as ProtoNotStartedPlayer, Response, ResponseType,
    StartedGamePerspective as ProtoStartedGamePerspective, StartedPlayer as ProtoStartedPlayer,
    UnknownCard as ProtoUnknownCard, UpdateGameResponse,
};

#[derive(Debug, Clone)]
pub struct Game {
    pub(crate) game_state: GameState,
}

impl Game {
    pub(crate) fn new() -> Self {
        Game {
            game_state: GameState::NotStarted(NotStartedGameState::new()),
        }
    }

    // TODO: Return result, NASTY race condition until you do. Also gross code.
    pub(crate) fn add_player(&mut self, player_name: String, id: String) {
        match self.game_state.clone() {
            GameState::NotStarted(mut old_game_state) => {
                old_game_state.players.push(NotStartedPlayer {
                    name: player_name,
                    id,
                    ready: false,
                });
                self.game_state = GameState::NotStarted(NotStartedGameState {
                    players: old_game_state.players,
                })
            }
            GameState::Started(_) => {
                println!("Tried to add a new player to a game that's already started.");
            }
        }
    }

    pub(crate) fn ready(&mut self, id: String) {
        match self.game_state.clone() {
            GameState::NotStarted(old_game_state) => {
                let mut players = old_game_state.players.clone();
                if let Some(player) = players.iter_mut().find(|p| p.id == id) {
                    println!("Found player. Previous status: {}", player.ready);
                    player.ready = !player.ready;
                }
                self.game_state = GameState::NotStarted(NotStartedGameState { players })
            }
            GameState::Started(_) => {
                println!("Tried to ready in a game that's already started. Ignoring message.")
            }
        }
    }

    pub(crate) fn start_game(&mut self) {
        match self.game_state.clone() {
            GameState::NotStarted(old_game_state) => {
                let deck = StartedGameState::new_deck();
                let (deck, players) = StartedGameState::deal(deck, &old_game_state.players);
                self.game_state = GameState::Started(StartedGameState {
                    players,
                    turn: 0,
                    deck,
                    discard_pile: Vec::new(),
                    hints: 8,
                    bombs: 3,
                    fireworks: StartedGameState::new_fireworks(),
                });
            }
            GameState::Started(_) => {
                println!("Tried to start a game that's already started. Ignoring message.")
            }
        }
    }
}

#[derive(Debug, Clone)]
pub enum GameState {
    NotStarted(NotStartedGameState),
    Started(StartedGameState),
}

#[derive(Debug, Clone)]
pub struct NotStartedGameState {
    pub(crate) players: Vec<NotStartedPlayer>,
}

impl NotStartedGameState {
    fn new() -> Self {
        NotStartedGameState {
            players: Vec::new(),
        }
    }
}

#[derive(Debug, Clone)]
pub struct NotStartedPlayer {
    name: String,
    id: String,
    pub(crate) ready: bool,
}

#[derive(Debug, Clone)]
pub struct StartedGameState {
    players: Vec<StartedPlayer>,
    turn: u8,
    deck: Vec<Card>,
    discard_pile: Vec<Card>,
    hints: u8,
    bombs: u8,
    fireworks: HashMap<CardColor, u8>,
}

#[derive(Debug, Clone)]
pub struct StartedPlayer {
    name: String,
    id: String,
    cards: [(Card, UnknownCard); 4],
}

impl StartedGameState {
    fn deal(
        mut deck: Vec<Card>,
        not_started_players: &Vec<NotStartedPlayer>,
    ) -> (Vec<Card>, Vec<StartedPlayer>) {
        let mut players = not_started_players
            .iter()
            .map(|nsp| {
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
                StartedPlayer {
                    name: nsp.name.clone(),
                    id: nsp.id.clone(),
                    cards,
                }
            })
            .collect();
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
pub struct GamePerspective {
    perspective: Perspective,
}

impl GamePerspective {
    pub fn from_state(game: &Game, player_index: u8) -> GamePerspective {
        match game.game_state.clone() {
            GameState::NotStarted(nsgs) => GamePerspective {
                perspective: Perspective::NotStarted(NotStartedGamePerspective::from_state(
                    &nsgs,
                    player_index,
                )),
            },
            GameState::Started(sgs) => GamePerspective {
                perspective: Perspective::Started(StartedGamePerspective::from_state(
                    &sgs,
                    player_index,
                )),
            },
        }
    }

    pub fn to_proto(&self) -> Response {
        let update_game_req = match &self.perspective {
            Perspective::NotStarted(nsgp) => UpdateGameResponse {
                started: false,
                game_perspective: Some(ProtoGamePerspective::NotStartedState(
                    ProtoNotStartedGamePerspective {
                        ready: nsgp.ready,
                        not_started_players: nsgp
                            .players
                            .iter()
                            .map(|c| c.clone().into())
                            .collect(),
                    },
                )),
            },
            Perspective::Started(sgs) => UpdateGameResponse {
                started: true,
                game_perspective: Some(ProtoGamePerspective::StartedState(
                    ProtoStartedGamePerspective {
                        my_hand: sgs.my_hand.iter().map(|c| c.clone().into()).collect(),
                        other_hands: sgs.other_hands.iter().map(|p| p.clone().into()).collect(),
                        turn: sgs.turn as i32,
                        deck: sgs.deck.iter().map(|c| c.clone().into()).collect(),
                        discard_pile: sgs.discard_pile.iter().map(|c| c.clone().into()).collect(),
                        hints: sgs.hints as i32,
                        bombs: sgs.bombs as i32,
                        fireworks: sgs
                            .fireworks
                            .iter()
                            .map(|(k, v)| (k.clone() as i32, *v as i32))
                            .collect(),
                    },
                )),
            },
        };
        Response {
            response_type: ResponseType::UpdateGame.into(),
            response: Some(generated::responses::response::Response::UpdateGame(
                update_game_req,
            )),
        }
    }
}

#[derive(Debug)]
pub enum Perspective {
    NotStarted(NotStartedGamePerspective),
    Started(StartedGamePerspective),
}

#[derive(Debug)]
pub struct NotStartedGamePerspective {
    ready: bool,
    players: Vec<NotStartedPlayer>,
}
impl NotStartedGamePerspective {
    pub fn from_state(game_state: &NotStartedGameState, player_index: u8) -> Self {
        let mut other_players = game_state.players.clone();
        let ready = other_players[player_index as usize].ready;
        other_players.remove(player_index as usize);
        other_players.rotate_left(player_index as usize);
        Self {
            ready,
            players: other_players,
        }
    }
}

#[derive(Debug)]
pub struct StartedGamePerspective {
    my_hand: [UnknownCard; 4],
    other_hands: Vec<StartedPlayer>,
    turn: u8,
    deck: Vec<Card>,
    discard_pile: Vec<Card>,
    hints: u8,
    bombs: u8,
    fireworks: HashMap<CardColor, u8>,
}

impl StartedGamePerspective {
    pub fn from_state(game_state: &StartedGameState, player_index: u8) -> Self {
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

        StartedGamePerspective {
            my_hand,
            other_hands,
            turn: game_state.turn,
            deck: game_state.deck.clone(),
            discard_pile: game_state.discard_pile.clone(),
            hints: game_state.hints,
            bombs: game_state.bombs,
            fireworks: game_state.fireworks.clone(),
        }
    }
}

impl From<NotStartedPlayer> for ProtoNotStartedPlayer {
    fn from(nsp: NotStartedPlayer) -> Self {
        ProtoNotStartedPlayer {
            name: nsp.name,
            id: nsp.id,
            ready: nsp.ready,
        }
    }
}
impl From<StartedPlayer> for ProtoStartedPlayer {
    fn from(sp: StartedPlayer) -> Self {
        ProtoStartedPlayer {
            name: sp.name,
            id: sp.id,
            cards: sp.cards.iter().map(|(c, _)| c.clone().into()).collect(),
            unknown_cards: sp.cards.iter().map(|(_, uc)| uc.clone().into()).collect(),
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
