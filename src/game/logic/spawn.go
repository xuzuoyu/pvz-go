package logic

import (
	"PvZ-go/src/config"
	"PvZ-go/src/entity"
	"PvZ-go/src/game/game_init"
)

func SpawnSystem(state *game_init.GameState) {
	for i := range state.Plants {
		p := &state.Plants[i]
		cfg := config.Plants[p.Type]

		p.Timer++
		if float64(p.Timer) >= cfg.Cooldown {
			state.Bullets = append(state.Bullets, entity.Bullet{
				Type:  cfg.Bullet,
				X:     float64(p.X),
				Y:     float64(p.Y),
				Alive: true,
			})

			p.Timer = 0
		}
	}
}
