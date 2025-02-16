package screens

import (
	"fmt"
	"image/color"

	"github.com/eric-ming2/hanabi/hanabi-frontend/state"
	"github.com/eric-ming2/hanabi/hanabi-frontend/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const readyButtonX, readyButtonY, readyButtonW, readyButtonH = 50, 200, 65, 30
const startButtonX, startButtonY, startButtonW, startButtonH = 350, 200, 65, 30

type NotStartedGameScreen struct {
	mouseDown bool
}

func NewNotStartedGameScreen() *NotStartedGameScreen {
	return &NotStartedGameScreen{
		mouseDown: false,
	}
}

func (s *NotStartedGameScreen) Update(workerReqChan chan<- websocket.WorkerRequest) error {
	// Handle button click
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !s.mouseDown {
			x, y := ebiten.CursorPosition()
			if x >= readyButtonX && x <= readyButtonX+readyButtonW && y >= readyButtonY && y <= readyButtonY+readyButtonH {
				workerReqChan <- websocket.WorkerRequest{
					Type: websocket.ReadyRequest,
				}
				fmt.Println("Ready button clicked!")
			} else if x >= startButtonX && x <= startButtonX+startButtonW && y >= startButtonY && y <= startButtonY+startButtonH {
				workerReqChan <- websocket.WorkerRequest{
					Type: websocket.StartGameRequest,
				}
				fmt.Println("Start button clicked!")
			}
		}
		s.mouseDown = true
	} else {
		s.mouseDown = false
	}
	return nil
}

func (s *NotStartedGameScreen) Draw(screen *ebiten.Image, username string, state *state.GameState) {
	if state == nil {
		return
	}

	// Draw ready button
	ebitenutil.DrawRect(screen, float64(readyButtonX), float64(readyButtonY), float64(readyButtonW), float64(readyButtonH), color.RGBA{200, 200, 200, 255})
	drawText(screen, "Ready", readyButtonX+10, readyButtonY+20, color.Black)

	// Draw start button
	ebitenutil.DrawRect(screen, float64(startButtonX), float64(startButtonY), float64(startButtonW), float64(startButtonH), color.RGBA{200, 200, 200, 255})
	drawText(screen, "Start", startButtonX, startButtonY+20, color.Black)

	drawUser(screen, 0, username, state.NotStartedState.Ready)
	for index, player := range state.NotStartedState.Players {
		drawUser(screen, index+1, player.Name, player.Ready)
	}
	// drawText(screen, fmt.Sprintf("%+v", *state), 0, 100, color.White)
}

type Coordinate struct {
	X int
	Y int
}

func drawUser(screen *ebiten.Image, playerIndex int, username string, ready bool) {
	usernameCoordinates := [4]Coordinate{
		{X: 300, Y: 300},
		{X: 10, Y: 150},
		{X: 500, Y: 150},
		{X: 300, Y: 10},
	}
	x := usernameCoordinates[playerIndex].X
	y := usernameCoordinates[playerIndex].Y
	drawText(screen, fmt.Sprintf("%+v", username), x, y, color.White)
	readyX := x + 10 + len(username)*7 // Adjust based on font size
	if ready {
		drawText(screen, fmt.Sprintf("Ready"), readyX, y, color.White)
	} else {
		drawText(screen, fmt.Sprintf("NotRe"), readyX, y, color.White)
	}
}
