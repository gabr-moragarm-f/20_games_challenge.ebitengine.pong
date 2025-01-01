package game

import (
	"github.com/gabr-moragarm-f/20_games_challenge.ebitengine.pong/game/phases/intro"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	currentPhase Phase
	settings     Settings
}

type Settings struct {
	width  int
	height int
}

const Ratio = 16.0 / 9.0

type Phase interface {
	Update() error
	Draw(screen *ebiten.Image)
	Layout(outsideWidth, outsideHeight int) (int, int)
}

func New() *Game {

	return &Game{
		currentPhase: intro.New(),
		settings: Settings{
			width:  1280,
			height: 960,
		},
	}
}

func (g *Game) Update() error {
	return g.currentPhase.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentPhase.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.currentPhase.Layout(g.settings.width, g.settings.height)
}

func (g *Game) SetPhase(phase Phase) {
	g.currentPhase = phase
}
