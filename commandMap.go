package main

import (
	"errors"
	"fmt"
)

func commandMapf(cfg *config, parameter string) error {
	response, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL, parameter)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = response.Next
	cfg.previousLocationsURL = response.Previous

	for _, val := range response.Results {
		fmt.Println(val.Name)
	}
	return nil
}

func commandMapb(cfg *config, parameter string) error {
	if cfg.previousLocationsURL == nil {
		return errors.New("You are on the first page.")
	}

	response, err := cfg.pokeapiClient.ListLocations(cfg.previousLocationsURL, parameter)
	if err != nil {
		return err
	}
	cfg.nextLocationsURL = response.Next
	cfg.previousLocationsURL = response.Previous

	for _, val := range response.Results {
		fmt.Println(val.Name)
	}
	return nil
}
