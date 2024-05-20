package main

import "fmt"

func pokedexCommand(cfg *config, _ ...string) error {
	fmt.Println("Your Pokemon:")
	for name := range cfg.myPokemon {
		fmt.Printf("- %s\n", name)
	}
	return nil
}
