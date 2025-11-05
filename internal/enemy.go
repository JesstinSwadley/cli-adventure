package internal

import "math/rand"

type Enemy struct {
	Name   string
	Health int32
}

var enemyNames = map[int]string{
	0: "Blackwing",
	1: "Wispmask",
	2: "Sparky",
	3: "Plaguebeast",
}

func (enm Enemy) NewEnemy() *Enemy {
	var enemy Enemy

	nameOption := rand.Intn(4)

	enemy.Name = enemyNames[nameOption]
	enemy.Health = rand.Int31n(10)

	return &enemy
}
