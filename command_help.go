package main

import (
	"fmt"
)

func commandHelp(cfg *config, args ...string) error {
	fmt.Println("")
	fmt.Println("Welcome to the Pokédex!")
	fmt.Println("Usage:")
	fmt.Println("")

	commands := getCommands()
	for _, c := range commands {
		fmt.Printf("%v: %v\n", c.name, c.description)
	}

	fmt.Println("")

	return nil
}
