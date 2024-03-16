package main

import "fmt"

func commandInspect(cfg *config, parameter string) error {

	pokemon, exists := (*cfg.pokedex)[parameter]
	if !exists {
		fmt.Println("You have not caught that pokemon.")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)

	}

	return nil
}
