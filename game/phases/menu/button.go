package menu

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Button struct {
	coordinates   func(screen *ebiten.Image) (int, int)
	x, y          int
	dimensions    func(screen *ebiten.Image) (int, int)
	width, height int
	Label         string
	OnClick       func(menu *Menu) error
	Hovered       bool
}

func (b *Button) Contains(mx, my int) bool {
	return mx >= b.x && mx <= b.x+b.width && my >= b.y && my <= b.y+b.height
}

func (b *Button) Update(menu *Menu, mx, my int, clicked bool) error {
	b.Hovered = b.Contains(mx, my)
	if b.Hovered && clicked {
		return b.OnClick(menu)
	}
	return nil
}

func (b *Button) Draw(screen *ebiten.Image) {
	var btnColor color.Color
	if b.Hovered {
		btnColor = color.RGBA{R: 255, G: 0, B: 0, A: 255} // Giallo quando è in hover
	} else {
		btnColor = color.RGBA{R: 0, G: 255, B: 0, A: 255} // Verde di default
	}

	// Disegna il rettangolo del pulsante
	ebitenutil.DrawRect(screen, float64(b.X), float64(b.Y), float64(b.Width), float64(b.Height), btnColor)

	// Disegna il testo del pulsante
	ebitenutil.DebugPrintAt(screen, b.Label, b.X+10, b.Y+10)
}
