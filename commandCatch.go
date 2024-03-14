package main

import "fmt"

func commandCatch(cfg *config, parameter string) error {
	response, err := cfg.pokeapiClient.CatchPokemon(parameter)
	if err != nil {
		return err
	}

	fmt.Println("Catched the following Pokemon:")
	fmt.Println(response.Name)

	return nil
}
