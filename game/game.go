package game

import (
	"github.com/gabr-moragarm-f/20_games_challenge.ebitengine.pong/game/phases/intro"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	currentPhase Phase
}

type Phase interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)
}

func New() *Game {

	return &Game{
		currentPhase: intro.New(),
	}
}

func (g *Game) Update() error {
	return g.currentPhase.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentPhase.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.currentPhase.Layout(outsideWidth, outsideHeight)
}

func (g *Game) SetPhase(phase Phase) {
	g.currentPhase = phase
}
