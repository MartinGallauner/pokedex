package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, parameter string) error {
	pokemon, err := cfg.pokeapiClient.CatchPokemon(parameter)
	if err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("Throwing a pokeball at %v...", parameter))

	chance := rand.Intn(400)

	if chance > pokemon.BaseExperience {
		fmt.Println(fmt.Sprintf("%v was caught!", pokemon.Name))
		(*cfg.pokedex)[pokemon.Name] = pokemon
		return nil
	}

	fmt.Println(fmt.Sprintf("%v escaped!", pokemon.Name))
	return nil
}
