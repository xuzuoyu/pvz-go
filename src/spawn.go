package src

func SpawnSystem(state *GameState) {
	for i := range state.Plants {
		p := &state.Plants[i]
		cfg := Plants[p.Type]

		p.Timer++
		if float64(p.Timer) >= cfg.Cooldown {
			state.Bullets = append(state.Bullets, Bullet{
				Type:  cfg.Bullet,
				X:     float64(p.X),
				Y:     float64(p.Y),
				Alive: true,
			})

			p.Timer = 0
		}
	}
}
