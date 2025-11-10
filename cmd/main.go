package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/JesstinSwadley/cli-adventure/internal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var player internal.Player

	p := player.NewPlayer(scanner)

	var enemy internal.Enemy

	enm := enemy.NewEnemy()

	fmt.Println("Your opponent is " + enm.Name + " and its health is " + strconv.Itoa(int(enm.Health)))

	// Use for loop to continuously perform actions
	turns := 0

	for turns != 1 {
		turns = Turn(p, enm, scanner)
	}
}

func Turn(p *internal.Player, enm *internal.Enemy, scanner *bufio.Scanner) int {
	// Need to add a system
	if p.Health > 0 && enm.Health > 0 {
		attack := p.Attack(scanner)

		enm.Health = enm.Health - int32(attack)

		fmt.Println(enm.Name + " current health is " + strconv.Itoa(int(enm.Health)))

		enmAttack := enm.Attack()

		p.Health = p.Health - int32(enmAttack)

		fmt.Println(enm.Name + " has attacked you, your health is now " + strconv.Itoa(int(p.Health)))

		return 0
	} else if p.Health <= 0 {
		fmt.Println("You have been defeated")

		return 1
	} else {
		fmt.Println("You have defeated " + enm.Name)

		return 1
	}
}
