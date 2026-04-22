package main

import (
	"PvZ-go/src"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := src.NewGame()

	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Mini PvZ - Go")

	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
