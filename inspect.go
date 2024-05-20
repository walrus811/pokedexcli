package main

import (
	"errors"
	"fmt"
)

func inspectCommand(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("no Pokemon specified. ex. catch pikachu")
	}

	name := params[0]

	pokemon, ok := cfg.myPokemon[name]

	if !ok {
		return errors.New("you have not caught that pokemon yet")
	}

	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("- %s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("- %s\n", t.Type.Name)
	}

	return nil
}
