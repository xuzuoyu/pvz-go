package entity

import (
	"PvZ-go/src/config"
)

type Plant struct {
	Type  config.PlantType
	X, Y  int
	Timer int
}
