package main

import "fmt"

func commandPokedex(config *config, args ...string) error {
	if len(config.caughtPokemon) == 0 {
		fmt.Println("Your Pokedex is empty.")
		return nil
	}

	fmt.Println("Your Pokedex:")
	for name := range config.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}
