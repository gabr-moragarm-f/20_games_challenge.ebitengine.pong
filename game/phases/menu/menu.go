package menu

import (
	"github.com/gabr-moragarm-f/20_games_challenge.ebitengine.pong/game/phases"
	"github.com/hajimehoshi/ebiten/v2"
)

const Identifier = phases.Menu

type Menu struct {
	buttons     []Button
	changePhase bool
	nextPhase   phases.Index
}

func New() *Menu {
	return &Menu{
		buttons: []Button{
			{
				x:      100,
				y:      100,
				width:  150,
				height: 50,
				Label:  "Start Game",
				OnClick: func(m *Menu) error {
					println("Start Game clicked")
					return nil
				},
			},
			{
				x:      100,
				y:      200,
				width:  150,
				height: 50,
				Label:  "Quit",
				OnClick: func(m *Menu) error {
					println("Quit clicked")
					return ebiten.Termination
				},
			},
		},
	}
}

func (m *Menu) Identify() phases.Index {
	return Identifier
}

func (m *Menu) Update() error {
	mx, my := ebiten.CursorPosition()
	clicked := ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft)

	for i := range m.buttons {
		if err := m.buttons[i].Update(m, mx, my, clicked); err != nil {
			return err
		}
	}

	return nil
}

func (m *Menu) Draw(screen *ebiten.Image) {
	for i := range m.buttons {
		m.buttons[i].Draw(screen)
	}
}

func (m *Menu) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (m *Menu) ChangePhase() (bool, phases.Index) {
	return m.changePhase, m.nextPhase
}
