package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListPokemons(locationName string) (Location, error) {
	url := baseURL + "/location-area/" + locationName

	if cacheEntry, cached := c.cache.Get(url); cached {
		locationResp := Location{}
		err := json.Unmarshal(cacheEntry, &locationResp)
		if err != nil {
			return Location{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	response, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer response.Body.Close()

	dat, err := io.ReadAll(response.Body)
	if err != nil {
		return Location{}, err
	}

	locationResp := Location{}
	err = json.Unmarshal(dat, &locationResp)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, dat)

	return locationResp, nil
}
