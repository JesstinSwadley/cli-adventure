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

	attack := p.Attack(scanner)

	enm.Health = enm.Health - int32(attack)

	if enm.Health > 0 {
		fmt.Println(enm.Name + " has survived your attack")
	} else {
		fmt.Println("You have defeated " + enm.Name)
	}
}
