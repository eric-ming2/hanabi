package screens

import (
	"github.com/eric-ming2/hanabi/hanabi-frontend/state"
	"github.com/eric-ming2/hanabi/hanabi-frontend/websocket"
	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
)

type Screen interface {
	Update(workerReqChan chan<- websocket.WorkerRequest) error
	Draw(state *state.GameState, screen *ebiten.Image)
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
