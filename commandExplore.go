package main

import "fmt"

func commandExplore(cfg *config, parameter string) error {
	response, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL, parameter)
	if err != nil {
		return err
	}

	for _, val := range response.Results {
		fmt.Println(val.Name)
	}
	return nil
}
