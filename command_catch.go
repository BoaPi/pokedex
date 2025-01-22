package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) < 1 {
		return errors.New("no Pokemon name given to catch")
	}

	if len(args) > 1 {
		return errors.New("too many Pokemon names given")
	}

	name := args[0]

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	basePossibility := 0.7
	difficulty := float64(pokemon.BaseExperience) / 400
	catchPosibility := basePossibility * (1 - difficulty)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	roll := r.Float64()

	if roll < catchPosibility {
		cfg.pokeapiClient.Registry[name] = pokemon
		fmt.Println(name, "was caught!")
	} else {
		fmt.Println(name, "escaped!")
	}

	return nil
}
