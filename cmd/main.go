package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/JesstinSwadley/cli-adventure/internal"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var player internal.Player

	p := player.NewPlayer(scanner)

	var enemy internal.Enemy

	enm := enemy.NewEnemy()

	fmt.Println(p.Name)
	fmt.Println(enm)
}
