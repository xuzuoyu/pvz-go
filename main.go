package main

import (
	ebiten2 "PvZ-go/src/game/ebiten"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := ebiten2.NewGame()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Mini PvZ - Go")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
