package screens

import (
	"fmt"
	"image/color"

	"github.com/eric-ming2/hanabi/hanabi-frontend/state"
	"github.com/eric-ming2/hanabi/hanabi-frontend/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

type StartedGameScreen struct{}

func NewStartedGameScreen() *StartedGameScreen {
	return &StartedGameScreen{}
}

func (s *StartedGameScreen) Update(workerReqChan chan<- websocket.WorkerRequest) error {
	return nil
}

func (s *StartedGameScreen) Draw(screen *ebiten.Image, username string, state *state.GameState) {
	if state != nil {
		drawText(screen, fmt.Sprintf("%+v", *state), 0, 100, color.White)
	}
}
