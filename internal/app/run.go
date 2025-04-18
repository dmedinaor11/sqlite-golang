package app

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Run() error {
	a := NewApp()

	if err := setUpDependencies(a); err != nil {
		return err
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome to the SQLite Go Version CLI!")
	for {
		fmt.Print("db> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		a.processInput(input)
	}
	return scanner.Err()
}
