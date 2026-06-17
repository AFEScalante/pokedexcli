package main

import (
	"fmt"

	"github.com/AFEScalante/pokedexcli/internal/pokeapi"
)

func NewPokedex() Pokedex {
	return Pokedex{
		data: make(map[string]pokeapi.Pokemon),
	}
}

func (p *Pokedex) Add(pokemon pokeapi.Pokemon) {
	p.data[pokemon.Name] = pokemon
}

func (p *Pokedex) PrintPokemons() {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range p.data {
		fmt.Printf("- %s\n", pokemon.Name)
	}
}

