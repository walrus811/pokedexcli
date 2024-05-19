package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("no location specified. ex. explore eterna-forest-area")
	}

	location := params[0]

	fmt.Printf("Exploring %s...\n", location)

	result, err := cfg.pokeClient.GetLocationDetails(location)

	if err != nil {
		return err
	}

	if result.PokemonEncounters == nil || len(result.PokemonEncounters) == 0 {
		return errors.New("no Pokemon encounters found")
	}

	fmt.Printf("Found Pokemon(%d):\n", len(result.PokemonEncounters))

	for _, pokemon := range result.PokemonEncounters {
		fmt.Printf("- %s\n", pokemon.Pokemon.Name)
	}

	return nil
}
