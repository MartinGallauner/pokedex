package main

import (
	"github.com/martingallauner/pokedex/internal/pokeapi"
	"time"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	pokedex := map[string]pokeapi.Pokemon{}
	cfg := &config{
		pokedex:       &pokedex,
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}
