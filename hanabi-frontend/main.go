package main

import (
	"fmt"
	"github.com/eric-ming2/hanabi/hanabi-frontend/screens"
	"github.com/eric-ming2/hanabi/hanabi-frontend/state"
	"github.com/eric-ming2/hanabi/hanabi-frontend/websocket"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	State         *state.GameState
	StartScreen   *screens.StartScreen
	GameScreen    *screens.GameScreen
	WorkerReqChan chan websocket.WorkerRequest
	WorkerResChan chan websocket.WorkerResponse
}

func (g *Game) Update() error {
	if g.State == nil {
		return g.StartScreen.Update(g.WorkerReqChan)
	}
	return g.GameScreen.Update(g.WorkerReqChan)
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.State == nil {
		g.StartScreen.Draw(g.State, screen)
	}
	g.GameScreen.Draw(g.State, screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 200
}

func main() {
	workerReqChan := make(chan websocket.WorkerRequest)
	workerResChan := make(chan websocket.WorkerResponse)
	game := &Game{
		StartScreen:   screens.NewStartScreen(),
		GameScreen:    screens.NewGameScreen(),
		WorkerReqChan: workerReqChan,
		WorkerResChan: workerResChan,
	}
	go websocket.ClientWorker(workerReqChan, workerResChan)
	go handleWorkerRes(workerResChan, game)
	ebiten.SetWindowSize(960, 540)
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
			fmt.Println("Your mom: {}", g.State)
		}
	}
}
