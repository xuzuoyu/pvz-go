package config

type PlantType string
type BulletType string
type ZombieType string

const (
	PlantPea  PlantType  = "pea"
	BulletPea BulletType = "pea"
)

type PlantConfig struct {
	HP       float64
	Cooldown float64
	Bullet   BulletType
}

type BulletConfig struct {
	Speed  float64
	Damage float64
	Alive  bool
}

var Plants = map[PlantType]PlantConfig{
	PlantPea: {
		HP:       100,
		Cooldown: 60,
		Bullet:   BulletPea,
	},
}

var Bullets = map[BulletType]BulletConfig{
	BulletPea: {
		Speed:  0.2,
		Damage: 5,
	},
}
