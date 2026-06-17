package main

import "fmt"

func commandInspect(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("you must provide a pokemon name")
	}
	pokemonName := args[0]

	pokemon, ok := cfg.pokedex.data[pokemonName]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	// General info
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	// Stats
	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %v\n", s.Stat.Name, s.BaseStat)
	}
	// Types
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  -%s\n", t.Type.Name)
	}

	return nil
}