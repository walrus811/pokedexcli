package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *config, params ...string) error {
	result, err := cfg.pokeClient.ListLocations(cfg.nextLocationAreaUrl)
	if err != nil {
		return err
	}
	cfg.nextLocationAreaUrl = result.Next
	cfg.prevLocationAreaUrl = result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapb(cfg *config, params ...string) error {
	if cfg.prevLocationAreaUrl == nil {
		cfg.nextLocationAreaUrl = nil
		return errors.New("no previous location area")
	}

	result, err := cfg.pokeClient.ListLocations(cfg.prevLocationAreaUrl)
	if err != nil {
		return err
	}

	cfg.nextLocationAreaUrl = result.Next
	cfg.prevLocationAreaUrl = result.Previous

	for _, location := range result.Results {
		fmt.Println(location.Name)
	}

	return nil
}
