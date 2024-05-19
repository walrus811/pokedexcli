package main

import (
	"time"

	"github.com/walrus811/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Second)
	cfg := &config{
		pokeClient: pokeClient,
	}
	startRepl(cfg)
}
