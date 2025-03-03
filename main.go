package main

import (
	"log"

	"github.com/gabr-moragarm-f/20_games_challenge.ebitengine.pong/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	ebiten.SetWindowSize(1280, 960)
	ebiten.SetWindowTitle("Pong - Ebitengine - 20 Games Challenge")
	// ebiten.SetFullscreen(true)
	// ebiten.SetWindowResizable(true)
	if err := ebiten.RunGame(game.New()); err != nil {
		log.Fatal(err)
	}
}
