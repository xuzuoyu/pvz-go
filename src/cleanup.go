package src

func CleanupSystem(state *GameState) {
	// Compact bullets in place and keep only active ones.
	bulletCount := 0
	for _, b := range state.Bullets {
		if b.Alive {
			state.Bullets[bulletCount] = b
			bulletCount++
		}
	}
	state.Bullets = state.Bullets[:bulletCount]

	// Compact zombies in place and keep only living ones.
	zombieCount := 0
	for _, z := range state.Zombies {
		if z.HP > 0 {
			state.Zombies[zombieCount] = z
			zombieCount++
		}
	}
	state.Zombies = state.Zombies[:zombieCount]
}
