package main

import (
	"time"

	pokeapi "github.com/BoaPi/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5*time.Second, 5*time.Minute)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startPepl(cfg)
}
