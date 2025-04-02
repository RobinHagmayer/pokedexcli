package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/RobinHagmayer/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*config) error
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationAreaResponse struct {
	Count    int            `json:"count"`
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
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

		cliCommands := getCliCommands()
		cmd, ok := cliCommands[cmdName]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}

		err := cmd.Callback(config)
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
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

func commandExit(config *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(config *config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()

	for _, cmd := range getCliCommands() {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func commandMapf(config *config) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if config.nextLocationURL != nil {
		url = *config.nextLocationURL
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Received non 200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apiResponse LocationAreaResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return err
	}

	config.nextLocationURL = apiResponse.Next
	config.previousLocationURL = apiResponse.Previous

	for i := range apiResponse.Results {
		fmt.Println(apiResponse.Results[i].Name)
	}

	return nil
}

func commandMapb(config *config) error {
	url := "https://pokeapi.co/api/v2/location-area"

	if config.previousLocationURL != nil {
		url = *config.previousLocationURL
	}

	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Received non 200 status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	var apiResponse LocationAreaResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		return err
	}

	config.nextLocationURL = apiResponse.Next
	config.previousLocationURL = apiResponse.Previous

	for i := range apiResponse.Results {
		fmt.Println(apiResponse.Results[i].Name)
	}

	return nil
}
