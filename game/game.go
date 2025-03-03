package game

import (
	"log"

	"github.com/gabr-moragarm-f/20_games_challenge.ebitengine.pong/game/phases"
	"github.com/gabr-moragarm-f/20_games_challenge.ebitengine.pong/game/phases/intro"
	"github.com/gabr-moragarm-f/20_games_challenge.ebitengine.pong/game/phases/menu"
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	currentPhase phases.Phase
	settings     Settings
}

type Settings struct {
	width  int
	height int
}

const Ratio = 16.0 / 9.0

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
	if err := g.currentPhase.Update(); err != nil {
		return err
	}

	if change, next := g.currentPhase.ChangePhase(); change {
		g.SetPhase(g.currentPhase.Identify(), next)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.currentPhase.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return g.currentPhase.Layout(g.settings.width, g.settings.height)
}

func (g *Game) SetPhase(from phases.Index, to phases.Index) {
	if !phases.ChangePermitted(from, to) {
		log.Fatalf("Phase change from %v to %v not permitted", from, to)
	}

	switch to {
	case phases.Intro:
		g.currentPhase = intro.New()
	case phases.Menu:
		g.currentPhase = menu.New()
	}
}
