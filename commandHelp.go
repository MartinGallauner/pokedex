package main

import "fmt"

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, value := range getCommands() {
		fmt.Println(value.name, ": ", value.description)
	}
	fmt.Println()
	return nil

}
