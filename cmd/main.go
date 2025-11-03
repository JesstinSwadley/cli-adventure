package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter your username: ")
	scanner.Scan()
	username := scanner.Text()
	fmt.Println(username)
}
