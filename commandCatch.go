package main

import "fmt"

func commandCatch(cfg *config, parameter string) error {
	pokemon, err := cfg.pokeapiClient.CatchPokemon(parameter)
	if err != nil {
		return err
	}

	fmt.Println("Catch the following Pokemon:")
	fmt.Println(pokemon.Name)

	//todo calculate chances of catching pokemon

	//todo add pokemon to pokedex

	(*cfg.pokedex)[pokemon.Name] = pokemon

	return nil
}
