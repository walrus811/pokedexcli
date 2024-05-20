package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, params ...string) error {
	if len(params) == 0 {
		return errors.New("no Pokemon specified. ex. catch pikachu")
	}

	pokemon := params[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)

	result, err := cfg.pokeClient.GetPokemon(pokemon)

	if err != nil {
		return err
	}

	baseExperience := result.BaseExperience

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	catchRate := random.Intn(baseExperience / 10 / 2)

	if catchRate <= 0 {
		fmt.Printf("%s was caught!\n", pokemon)
	} else {
		fmt.Printf("%s broke free!\n", pokemon)
		return nil
	}

	cfg.myPokemon[pokemon] = result

	return nil
}
