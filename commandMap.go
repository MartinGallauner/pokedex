package main

import (
	"fmt"
	"github.com/martingallauner/pokedex/internal/pokeapi"
)

func commandMapf() error {
	response := pokeapi.GetLocations()
	for _, val := range response.Results {
		fmt.Println(val.Name)
	}
	return nil
}

func commandMapb() error {
	response := pokeapi.GetLocations()
	for _, val := range response.Results {
		fmt.Println(val.Name)
	}
	return nil
}
