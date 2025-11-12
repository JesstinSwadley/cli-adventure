package internal

import (
	"bufio"
	"fmt"
	"math/rand"
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

func (p Player) Attack(scanner *bufio.Scanner) int {
	fmt.Println("Do you want to attack yes/no?")
	scanner.Scan()
	response := scanner.Text()

	if response == "no" {
		return 0
	} else {
		return 5
	}
}

func (p Player) Heal() int {
	fmt.Println("Healing")

	healing := rand.Intn(5) + 1
	return healing
}
