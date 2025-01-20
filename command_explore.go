package main

import (
	"fmt"
)

func commandExplore(cfg *config, param *string) error {
	pokemonResp, err := cfg.pokeapiClient.ListPokemons(param)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %s...\n", *param)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range pokemonResp.PokemonEncounters {
		fmt.Println("- ", pokemon.Pokemon.Name)
	}

	return nil
}
