package main

import (
	"fmt"

	"github.com/eric-ming2/hanabi/hanabi-frontend/screens"
	"github.com/eric-ming2/hanabi/hanabi-frontend/state"
	"github.com/eric-ming2/hanabi/hanabi-frontend/websocket"
	"github.com/google/uuid"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	username             string
	id                   string
	State                *state.GameState
	StartScreen          *screens.StartScreen
	NotStartedGameScreen *screens.NotStartedGameScreen
	StartedGameScreen    *screens.StartedGameScreen
	WorkerReqChan        chan websocket.WorkerRequest
	WorkerResChan        chan websocket.WorkerResponse
}

func (g *Game) Update() error {
	if g.State == nil {
		return g.StartScreen.Update(g.WorkerReqChan, &g.username, g.id)
	} else if g.State.Started {
		return g.StartedGameScreen.Update(g.WorkerReqChan)
	} else {
		return g.NotStartedGameScreen.Update(g.WorkerReqChan)
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.State == nil {
		g.StartScreen.Draw(screen, &g.username)
	} else if g.State.Started {
		g.StartedGameScreen.Draw(screen, g.username, g.State)
	} else {
		g.NotStartedGameScreen.Draw(screen, g.username, g.State)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 360
}

func main() {
	workerReqChan := make(chan websocket.WorkerRequest)
	workerResChan := make(chan websocket.WorkerResponse)
	game := &Game{
		id:                   uuid.New().String(),
		StartScreen:          screens.NewStartScreen(),
		NotStartedGameScreen: screens.NewNotStartedGameScreen(),
		StartedGameScreen:    screens.NewStartedGameScreen(),
		WorkerReqChan:        workerReqChan,
		WorkerResChan:        workerResChan,
	}
	go websocket.ClientWorker(workerReqChan, workerResChan, game.id)
	go handleWorkerRes(workerResChan, game)
	ebiten.SetWindowSize(1280, 720)
	ebiten.SetWindowTitle("Hanabi")
	if err := ebiten.RunGame(game); err != nil {
		fmt.Println(err)
	}
}

func handleWorkerRes(ch <-chan websocket.WorkerResponse, g *Game) {
	for res := range ch {
		switch res.Type {
		case websocket.ConnectFailed:
			fmt.Println("Unimplemented atm")
		case websocket.UpdateGameState:
			g.State = res.Payload.(*state.GameState)
			fmt.Printf("Update Game State: %+v\n", g.State)
		}
	}
}
