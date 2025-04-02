package main

import (
	"time"

	"github.com/RobinHagmayer/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startREPL(cfg)
}
