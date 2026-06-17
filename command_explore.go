package main

import "fmt"

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide an area name")
	}
	area_name := args[0]
	locationData, err := cfg.pokeapiClient.GetLocation(area_name)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", area_name)
	for _, pokemon := range locationData.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
