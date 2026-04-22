package entity

import (
	"PvZ-go/src/config"
)

type Bullet struct {
	Type  config.BulletType
	X, Y  float64
	Alive bool
}
