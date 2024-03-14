package main

import (
	"bufio"
	"fmt"
	"github.com/martingallauner/pokedex/internal/pokeapi"
	"github.com/martingallauner/pokedex/internal/pokecache"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, string) error
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		var parameter string
		if len(words) == 2 {
			parameter = words[1]
		}

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, parameter)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Shows the next page of locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "Shows the previous page of locations",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore",
			description: "Pass the name of the location you want to explore",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Pass the name of the pokemon you want to try to catch",
			callback:    commandCatch,
		},
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type config struct {
	pokeapiClient        pokeapi.Client
	pokeCache            pokecache.Cache
	nextLocationsURL     *string
	previousLocationsURL *string
}
