package src

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	state *GameState
}

func NewGame() *Game {
	return &Game{
		state: &GameState{
			Plants: []Plant{
				{Type: "pea",
					X: 2,
					Y: 2},
			},
			Zombies: []Zombie{
				{
					X:     10,
					Y:     10,
					HP:    20,
					Speed: 0.05,
				},
			},
		},
	}
}

func (g *Game) Update() error {
	SpawnSystem(g.state)
	ZombieMoveSystem(g.state)
	HitSystem(g.state)
	CleanupSystem(g.state)

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	// 清屏（必须）
	screen.Fill(color.Black)

	// 植物
	for _, p := range g.state.Plants {
		ebitenutil.DrawRect(screen,
			float64(p.X*50),
			float64(p.Y*50),
			40, 40,
			color.RGBA{0, 255, 0, 255})
	}

	// 僵尸
	for _, z := range g.state.Zombies {
		ebitenutil.DrawRect(screen,
			float64(z.X*50),
			float64(z.Y*50),
			40, 40,
			color.RGBA{255, 0, 0, 255})
	}

	// 子弹
	for _, b := range g.state.Bullets {
		ebitenutil.DrawRect(screen,
			b.X*50,
			float64(b.Y*50),
			10, 10,
			color.RGBA{255, 255, 0, 255})
	}
}

func (g *Game) Layout(w, h int) (int, int) {
	return 800, 600
}
