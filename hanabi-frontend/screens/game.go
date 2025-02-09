package screens

import (
	"fmt"
	"github.com/eric-ming2/hanabi/hanabi-frontend/state"
	"github.com/eric-ming2/hanabi/hanabi-frontend/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
)

type GameScreen struct{}

func NewGameScreen() *GameScreen {
	return &GameScreen{}
}

func (s *GameScreen) Update(workerReqChan chan<- websocket.WorkerRequest) error {
	return nil
}

func (s *GameScreen) Draw(state *state.GameState, screen *ebiten.Image) {
	if state != nil {
		drawText(screen, fmt.Sprintf("%+v", *state), 0, 100, color.White)
	}
}
