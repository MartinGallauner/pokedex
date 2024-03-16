package main

import "fmt"

func commandInspect(cfg *config, parameter string) error {

	pokemon, exists := (*cfg.pokedex)[parameter]
	if !exists {
		fmt.Println("You have not caught that pokemon.")
		return nil
	}
	fmt.Println(pokemon)
	return nil
}
