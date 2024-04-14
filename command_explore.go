package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no locations were provided")
	}

	for _, locationAreaName := range args {
		locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
		if err != nil {
			return err
		}
		fmt.Printf("Exploring %s...\n", locationArea.Name)
		for _, pokemon := range locationArea.PokemonEncounters {
			fmt.Println(pokemon.Pokemon.Name)
		}
	}

	return nil
}
