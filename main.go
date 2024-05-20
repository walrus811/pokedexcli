package main

import (
	"time"

	"github.com/walrus811/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	myPokemon := make(map[string]pokeapi.Pokemon)
	cfg := &config{
		pokeClient: pokeClient,
		myPokemon:  myPokemon,
	}
	startRepl(cfg)
}
