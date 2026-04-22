package src

func BulletMoveSystem(state *GameState) {
	for i := range state.Bullets {
		b := &state.Bullets[i]
		if !b.Alive {
			continue
		}

		cfg := Bullets[b.Type]
		b.X += cfg.Speed

		if b.X > 20 {
			b.Alive = false
		}
	}
}

func ZombieMoveSystem(state *GameState) {
	for i := range state.Zombies {
		state.Zombies[i].X -= state.Zombies[i].Speed
	}
}
