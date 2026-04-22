package src

import (
	"math"
)

func HitSystem(state *GameState) {
	for bi := range state.Bullets {
		b := &state.Bullets[bi]

		if !b.Alive {
			continue
		}

		for zi := range state.Zombies {
			z := &state.Zombies[zi]

			if z.HP <= 0 {
				continue
			}

			if math.Abs(b.X-z.X) < 0.3 && b.Y == z.Y {
				cfg := Bullets[b.Type]
				z.HP -= cfg.Damage
				b.Alive = false
				break
			}
		}
	}

}
