package main

import "fmt"

func commandExplore(cfg *config, parameter string) error {
	response, err := cfg.pokeapiClient.ExploreLocation(cfg.nextLocationsURL, parameter)
	if err != nil {
		return err
	}

	if len(response.PokemonEncounters) == 0 {
		return nil
	}

	fmt.Println("Found the following Pokemon:")

	for _, val := range response.PokemonEncounters {
		fmt.Println(fmt.Sprintf("- %v", val.Pokemon.Name))
	}
	return nil
}
