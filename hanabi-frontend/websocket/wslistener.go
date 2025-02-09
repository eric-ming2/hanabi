package websocket

import (
	"github.com/eric-ming2/hanabi/hanabi-frontend/generated"
	"github.com/eric-ming2/hanabi/hanabi-frontend/state"
	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"log"
)

func listen(conn *websocket.Conn, wsResChan chan WorkerResponse) {
	for {
		messageType, message, err := conn.ReadMessage()
		log.Printf("Just read a message!")
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		// Handle different message types
		switch messageType {
		case websocket.TextMessage:
			log.Printf("Received text message: %s", message)
		case websocket.BinaryMessage:
			var response generated.Response
			err := proto.Unmarshal(message, &response)
			if err != nil {
				log.Printf("Unable to unmarshal proto: %v", err)
				return
			}
			log.Printf("Ur mom: {}", response.ResponseType)
			switch response.ResponseType {
			case generated.ResponseType_UPDATE_GAME:
				log.Printf("Received UPDATE_GAME message: %v", response.GetUpdateGame())
				wsResChan <- WorkerResponse{
					Type:    UpdateGameState,
					Payload: parseUpdateGame(response.GetUpdateGame()),
				}
			default:
				log.Printf("Unrecognized response from server.")
			}
		default:
			log.Printf("Received unknown message type (%d): %x", messageType, message)
		}
	}
}

func parseUpdateGame(updateGameProto *generated.UpdateGameResponse) *state.GameState {
	started := updateGameProto.Started
	if !started {
		var players []state.NotStartedPlayer
		for _, p := range updateGameProto.GetNotStartedState().GetNotStartedPlayers() {
			players = append(players, state.NotStartedPlayer{
				Name:  p.Name,
				Id:    p.Id,
				Ready: false,
			})
		}
		return &state.GameState{
			Started: started,
			NotStartedState: state.NotStartedGameState{
				Players: players,
			},
		}
	} else {
		var myHand []state.UnknownCard
		for _, uc := range updateGameProto.GetStartedState().GetMyHand() {
			myHand = append(myHand, parseUnknownCard(uc.Color, uc.Num))
		}
		var otherPlayers []state.StartedPlayer
		for _, oh := range updateGameProto.GetStartedState().GetOtherHands() {
			var cards []state.Card
			for _, c := range oh.Cards {
				cards = append(cards, state.Card{
					Color: state.CardColor(c.Color),
					Num:   uint8(c.Num),
				})
			}
			var unknownCards []state.UnknownCard
			for _, uc := range oh.UnknownCards {
				unknownCards = append(unknownCards, parseUnknownCard(uc.Color, uc.Num))
			}
			otherPlayers = append(otherPlayers, state.StartedPlayer{
				Name:         oh.Name,
				Id:           oh.Id,
				Cards:        cards,
				UnknownCards: unknownCards,
			})
		}
		var deck []state.Card
		for _, c := range updateGameProto.GetStartedState().GetDeck() {
			deck = append(deck, state.Card{
				Color: state.CardColor(c.Color),
				Num:   uint8(c.Num),
			})
		}
		var discardPile []state.Card
		for _, c := range updateGameProto.GetStartedState().GetDiscardPile() {
			discardPile = append(discardPile, state.Card{
				Color: state.CardColor(c.Color),
				Num:   uint8(c.Num),
			})
		}
		fireworks := make(map[state.CardColor]uint8)
		for k, v := range updateGameProto.GetStartedState().GetFireworks() {
			fireworks[state.CardColor(k)] = uint8(v)
		}
		return &state.GameState{
			Started: true,
			StartedState: state.StartedGameState{
				MyHand:       myHand,
				OtherPlayers: otherPlayers,
				Turn:         uint8(updateGameProto.GetStartedState().Turn),
				Deck:         deck,
				DiscardPile:  discardPile,
				Hints:        uint8(updateGameProto.GetStartedState().Hints),
				Bombs:        uint8(updateGameProto.GetStartedState().Bombs),
				Fireworks:    fireworks,
			},
		}
	}
}

func parseUnknownColor(c *generated.CardColor) (bool, state.CardColor) {
	if c == nil {
		return false, state.White
	} else {
		return true, state.CardColor(*c)
	}
}

func parseUnknownNum(n *int32) (bool, uint8) {
	if n == nil {
		return false, 0
	} else {
		return true, uint8(*n)
	}
}

func parseUnknownCard(c *generated.CardColor, n *int32) state.UnknownCard {
	colorKnown, color := parseUnknownColor(c)
	numKnown, num := parseUnknownNum(n)
	return state.UnknownCard{
		ColorKnown: colorKnown,
		Color:      color,
		NumKnown:   numKnown,
		Num:        num,
	}
}
