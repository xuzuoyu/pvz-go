package logic

import (
	"PvZ-go/src/config"
	"PvZ-go/src/game/game_init"
)

func BulletMoveSystem(state *game_init.GameState) {
	for i := range state.Bullets {
		b := &state.Bullets[i]
		if !b.Alive {
			continue
		}

		cfg := config.Bullets[b.Type]
		b.X += cfg.Speed

		if b.X > 20 {
			b.Alive = false
		}
	}
}

func ZombieMoveSystem(state *game_init.GameState) {
	for i := range state.Zombies {
		state.Zombies[i].X -= state.Zombies[i].Speed
	}
}
