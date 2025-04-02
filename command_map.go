package main

import (
	"errors"
	"fmt"
)

func commandMapf(config *config) error {
	locationResp, err := config.pokeapiClient.ListLocations(config.nextLocationURL)
	if err != nil {
		return err
	}

	config.nextLocationURL = locationResp.Next
	config.previousLocationURL = locationResp.Previous

	for i := range locationResp.Results {
		fmt.Println(locationResp.Results[i].Name)
	}

	return nil
}

func commandMapb(config *config) error {
	if config.previousLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := config.pokeapiClient.ListLocations(config.previousLocationURL)
	if err != nil {
		return err
	}

	config.nextLocationURL = locationResp.Next
	config.previousLocationURL = locationResp.Previous

	for i := range locationResp.Results {
		fmt.Println(locationResp.Results[i].Name)
	}

	return nil
}
