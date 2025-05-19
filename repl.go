package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/RobinHagmayer/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config, ...string) error
}

func startREPL(config *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		cmdName := input[0]
		args := []string{}
		if len(input) > 1 {
			args = input[1:]
		}

		cmd, ok := getCliCommands()[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.Callback(config, args...)
		if err != nil {
			errMsg := fmt.Errorf("Error after calling %s: %w", cmd.Name, err)
			fmt.Println(errMsg)
			continue
		}
	}
}

func getCliCommands() map[string]CliCommand {
	return map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
			Callback:    commandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    commandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays the next 20 location areas",
			Callback:    commandMapf,
		},
		"mapb": {
			Name:        "mapb",
			Description: "Displays the previous 20 location areas",
			Callback:    commandMapb,
		},
		"explore": {
			Name:        "explore <location_name>",
			Description: "List all Pokemons from the given location",
			Callback:    commandExplore,
		},
		"catch": {
			Name:        "catch <pokemon_name>",
			Description: "Try to catch a Pokemon and add it to your Pokedex",
			Callback:    commandCatch,
		},
		"inspect": {
			Name:        "inspect <pokemon_name>",
			Description: "Inspect a caught Pokemon",
			Callback:    commandInspect,
		},
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}
