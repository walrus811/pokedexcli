package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/walrus811/pokedexcli/internal/pokeapi"
	"github.com/walrus811/pokedexcli/internal/pokecache"
)

type config struct {
	pokeClient          pokeapi.Client
	pokeCache           *pokecache.Cache
	nextLocationAreaUrl *string
	prevLocationAreaUrl *string
}

var cliCommands = make(map[string]cliCommand)

func startRepl(cfg *config) {
	initCommands()

	scanner := bufio.NewScanner(os.Stdin)
	commands := cliCommands

	for {
		fmt.Print("pokedex > ")
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "can't take input like that.", err)
		}

		scanned := scanner.Text()
		splitted := strings.Fields(scanned)

		command, ok := commands[splitted[0]]

		if !ok {
			fmt.Printf("'%s' is not a valid command. check the 'help' command!\n", scanned)
			continue
		}

		err := command.callback(cfg, splitted[1:]...)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func initCommands() {
	cliCommands["help"] = cliCommand{
		name:        "help",
		description: "Displays a help message",
		callback:    commandHelp,
	}
	cliCommands["exit"] = cliCommand{
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    commandExit,
	}
	cliCommands["map"] = cliCommand{
		name:        "map",
		description: "see next 20 location areas of Pokemon world",
		callback:    commandMap,
	}
	cliCommands["mapb"] = cliCommand{
		name:        "mapb",
		description: "sett previous 20 location areas of Pokemon world",
		callback:    commandMapb,
	}
	cliCommands["explore"] = cliCommand{
		name:        "explore",
		description: "explore the given location area",
		callback:    commandExplore,
	}
	cliCommands["catch"] = cliCommand{
		name:        "catch",
		description: "I'm gonna catch them all! yeah!! mezase Pokemon master!`",
		callback:    commandExplore,
	}
}
