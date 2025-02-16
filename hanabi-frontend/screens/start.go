package screens

import (
	"fmt"
	"image/color"

	"github.com/eric-ming2/hanabi/hanabi-frontend/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type StartScreen struct {
	Id             string
	CursorBlink    bool
	ConnectPressed bool
	TickCount      int
}

func NewStartScreen() *StartScreen {
	return &StartScreen{}
}

const connectButtonX = 50
const connectButtonY = 160
const connectButtonW = 65
const connectButtonH = 30

func (s *StartScreen) Update(workerReqChan chan<- websocket.WorkerRequest, username *string, id string) error {
	// Handle input characters
	for _, char := range ebiten.InputChars() {
		if char >= 32 && char <= 126 { // Printable ASCII characters
			*username += string(char)
		}
	}

	// Handle backspace
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) && len(*username) > 0 {
		*username = (*username)[:len(*username)-1]
	}

	// Update cursor blink state
	s.TickCount++
	if s.TickCount%30 == 0 { // Blink every 30 frames
		s.CursorBlink = !s.CursorBlink
	}

	// Handle button click
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= connectButtonX && x <= connectButtonX+connectButtonW && y >= connectButtonY && y <= connectButtonY+connectButtonH {
			if !s.ConnectPressed {
				s.ConnectPressed = true
				workerReqChan <- websocket.WorkerRequest{
					Type: websocket.ConnectRequest,
					Payload: websocket.ConnectRequestPayload{
						Id:       id,
						Username: *username,
					},
				}
			}
			fmt.Println("Button clicked!")
		}
	}
	return nil
}

func (s *StartScreen) Draw(screen *ebiten.Image, username *string) {
	drawText(screen, "Welcome To Hanabi", 125, 30, color.White)
	drawText(screen, "Username:", 50, 90, color.White)
	// Draw input box
	boxX, boxY, boxW, boxH := 50, 100, 300, 50
	ebitenutil.DrawRect(screen, float64(boxX), float64(boxY), float64(boxW), float64(boxH), color.RGBA{200, 200, 200, 255})
	ebitenutil.DrawRect(screen, float64(boxX+2), float64(boxY+2), float64(boxW-4), float64(boxH-4), color.Black)

	// Draw input text
	drawText(screen, *username, boxX+10, boxY+30, color.White)

	// Draw blinking cursor
	if s.CursorBlink {
		cursorX := boxX + 10 + len(*username)*7 // Adjust based on font size
		ebitenutil.DrawRect(screen, float64(cursorX), float64(boxY+10), 2, float64(boxH-20), color.White)
	}
	// Draw connect button
	buttonX, buttonY, buttonW, buttonH := 50, 160, 65, 30
	ebitenutil.DrawRect(screen, float64(buttonX), float64(buttonY), float64(buttonW), float64(buttonH), color.RGBA{200, 200, 200, 255})
	drawText(screen, "Connect", buttonX+10, buttonY+20, color.Black)
}
