package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon in Pokedex:")
	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Height: %d\n", pokemon.Weight)
		fmt.Println("Types:")
		for _, typ := range pokemon.Types {
			fmt.Printf(" - %s ", typ.Type.Name)
		}
		fmt.Println()
	}

	return nil
}
