package main

import (
	"errors"
	"fmt"
)

func commandMapF(cfg *config, args ...string) error {
	locationsData, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsData.Next
	cfg.prevLocationURL = locationsData.Previous

	for _, location := range locationsData.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func commandMapB(cfg *config, args ...string) error {
	if cfg.prevLocationURL == nil {
		return errors.New("you're on the first page")
	}

	locationsData, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationURL)
	if err != nil {
		return err
	}

	cfg.nextLocationURL = locationsData.Next
	cfg.prevLocationURL = locationsData.Previous

	for _, loc := range locationsData.Results {
		fmt.Println(loc.Name)
	}
	return nil
}
