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

	fmt.Println(p.Name)
}
