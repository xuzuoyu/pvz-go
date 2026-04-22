package game_init

import "PvZ-go/src/entity"

type GameState struct {
	Plants  []entity.Plant
	Zombies []entity.Zombie
	Bullets []entity.Bullet
}
