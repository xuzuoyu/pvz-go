package src

func CleanupSystem(state GameState) {
	// bullet
	n := 0
	for _, b := range Bullets {
		if b.Alive {
			state.Bullets[n] = b
			n++
		}
	}

	// zombie
	n = 0
	for _, z := range state.Zombies {
		if z.HP > 0 {
			state.Zombies[n] = z
			n++
		}
	}
	state.Zombies = state.Zombies[:n]
}
