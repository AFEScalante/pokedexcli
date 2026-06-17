package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/AFEScalante/pokedexcli/internal/pokeapi"
)

type Pokedex struct {
	data map[string]pokeapi.Pokemon
}

type config struct {
	pokeapiClient pokeapi.Client
	pokedex Pokedex
	nextLocationURL *string
	prevLocationURL *string
}

func startREPL(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)

	commands := getCommands()

	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		cmd, exists := commands[commandName]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
		}
	}
}

func cleanInput(text string) []string {
	return strings.Fields(strings.ToLower(text))
}

type cliCommand struct {
	name string
	description string
	callback func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name: "exit",
			description: "Exit the Pokedex",
			callback: commandExit,
		},
		"map": {
			name: "map",
			description: "Get the next page of locations",
			callback: commandMapF,
		},
		"mapb": {
			name: "mapb",
			description: "Get the previous page of locations",
			callback: commandMapB,
		},
		"explore": {
			name: "explore",
			description: "Explore a specific <area_name>",
			callback: commandExplore,
		},
		"catch": {
			name: "catch",
			description: "Catch a Pokemon by <pokemon_name>",
			callback: commandCatch,
		},
		"inspect": {
			name: "inspect",
			description: "Inspect pokemon availables in Pokedex",
			callback: commandInspect,
		},
		"pokedex": {
			name: "pokedex",
			description: "Print your pokedex",
			callback: commandPokedex,
		},
		"help": {
			name: "help",
			description: "Displays a help message",
			callback: commandHelp,
		},
	}
	return commands
}
