package main

import "fmt"

func commandPokedex(cfg *config, parameter string) error {

	fmt.Println("Your Pokedex:")
	for _, pokemon := range *cfg.pokedex {
		fmt.Printf(" - %v\n", pokemon.Name)
	}
	return nil
}
