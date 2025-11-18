package internal

import (
	"bufio"
	"fmt"
	"strconv"
)

func Actions(scanner *bufio.Scanner) {
	fmt.Println("What action would you like to do?")
	fmt.Println("1. Heal")
	fmt.Println("2. Attack")

	scanner.Scan()

	reader := scanner.Text()
	choice, err := strconv.Atoi(reader)

	if err != nil {
		panic(err)
	}

	switch choice {
	case 1:
		fmt.Println("You choice heal")
	case 2:
		fmt.Println("You choice attack")
	}
}
