package internal

import (
	"bufio"
	"fmt"
	"strconv"
)

type Player struct {
	Name   string
	Health int32
}

func (p Player) NewPlayer(scanner *bufio.Scanner) *Player {
	var player Player

	fmt.Print("Please enter your player name: ")
	scanner.Scan()
	player.Name = scanner.Text()

	fmt.Print("Total Player Health: ")
	scanner.Scan()
	health, err := strconv.ParseInt(scanner.Text(), 10, 32)

	player.Health = int32(health)

	if err != nil {
		panic(err)
	}

	return &player
}
