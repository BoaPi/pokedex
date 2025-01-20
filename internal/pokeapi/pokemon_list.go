package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(location *string) (RespPokemon, error) {
	url := baseURL + "/location-area/"
	if *location == "" {
		return RespPokemon{}, errors.New("No location to explore given")
	}

	url = url + *location

	if cacheEntry, cached := c.cache.Get(url); cached {
		pokemonResp := RespPokemon{}
		err := json.Unmarshal(cacheEntry, &pokemonResp)
		if err != nil {
			return RespPokemon{}, err
		}
		return pokemonResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespPokemon{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return RespPokemon{}, err
	}
	defer response.Body.Close()

	dat, err := io.ReadAll(response.Body)
	if err != nil {
		return RespPokemon{}, err
	}

	pokemonResp := RespPokemon{}
	err = json.Unmarshal(dat, &pokemonResp)
	if err != nil {
		return RespPokemon{}, err
	}

	c.cache.Add(url, dat)
	return pokemonResp, nil
}
