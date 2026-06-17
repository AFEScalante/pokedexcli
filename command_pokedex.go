package main

func commandPokedex(cfg *config, args ...string) error {
	cfg.pokedex.PrintPokemons()
	return nil
}