package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type Game struct {
	inputText      string
	cursorBlink    bool
	connectPressed bool
	tickCount      int
}

const connectButtonX = 50
const connectButtonY = 160
const connectButtonW = 65
const connectButtonH = 30

func (g *Game) Update() error {
	// Handle input characters
	for _, char := range ebiten.InputChars() {
		if char >= 32 && char <= 126 { // Printable ASCII characters
			g.inputText += string(char)
		}
	}

	// Handle backspace
	if ebiten.IsKeyPressed(ebiten.KeyBackspace) && len(g.inputText) > 0 {
		g.inputText = g.inputText[:len(g.inputText)-1]
	}

	// Update cursor blink state
	g.tickCount++
	if g.tickCount%30 == 0 { // Blink every 30 frames
		g.cursorBlink = !g.cursorBlink
	}

	// Handle button click
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		if x >= connectButtonX && x <= connectButtonX+connectButtonW && y >= connectButtonY && y <= connectButtonY+connectButtonH {
			g.connectPressed = true
			fmt.Println("Button clicked!")
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawText(screen, "Welcome To Hanabi", 125, 30, color.White)
	drawText(screen, "Username:", 50, 90, color.White)
	// Draw input box
	boxX, boxY, boxW, boxH := 50, 100, 300, 50
	ebitenutil.DrawRect(screen, float64(boxX), float64(boxY), float64(boxW), float64(boxH), color.RGBA{200, 200, 200, 255})
	ebitenutil.DrawRect(screen, float64(boxX+2), float64(boxY+2), float64(boxW-4), float64(boxH-4), color.Black)

	// Draw input text
	drawText(screen, g.inputText, boxX+10, boxY+30, color.White)

	// Draw blinking cursor
	if g.cursorBlink {
		cursorX := boxX + 10 + len(g.inputText)*7 // Adjust based on font size
		ebitenutil.DrawRect(screen, float64(cursorX), float64(boxY+10), 2, float64(boxH-20), color.White)
	}
	// Draw connect button
	buttonX, buttonY, buttonW, buttonH := 50, 160, 65, 30
	ebitenutil.DrawRect(screen, float64(buttonX), float64(buttonY), float64(buttonW), float64(buttonH), color.RGBA{200, 200, 200, 255})
	drawText(screen, "Connect", buttonX+10, buttonY+20, color.Black)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 400, 200
}

// Helper to draw text
func drawText(dst *ebiten.Image, text string, x, y int, clr color.Color) {
	face := basicfont.Face7x13
	d := &font.Drawer{
		Dst:  dst,
		Src:  image.NewUniform(clr),
		Face: face,
		Dot:  fixed.P(x, y),
	}
	d.DrawString(text)
}

func main() {
	ebiten.SetWindowSize(960, 540)
	ebiten.SetWindowTitle("Hanabi")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		fmt.Println(err)
	}
}
