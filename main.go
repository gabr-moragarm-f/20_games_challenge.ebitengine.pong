package main

import (
	"log"

	"github.com/gabr-moragarm-f/20_games_challenge.ebitengine.pong/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := game.New()

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Pong - Ebitengine - 20 Games Challenge")
	if err := ebiten.RunGame(game.New()); err != nil {
		log.Fatal(err)
	}
}
