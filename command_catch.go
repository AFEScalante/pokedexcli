package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}
	pokemonName := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
	if rand.Intn(pokemon.BaseExperience) <= 50 {
		fmt.Printf("%s was caught!\n", pokemonName)
		cfg.pokedex.Add(pokemon)
		return nil
	}

	fmt.Printf("%s escaped!\n", pokemonName)

	return nil
}
