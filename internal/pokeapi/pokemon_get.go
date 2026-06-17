package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	cachedData, cacheExists := c.pokeCache.Get(url)
	if cacheExists {
		pokemonRes := Pokemon{}
		if err := json.Unmarshal(cachedData, &pokemonRes); err != nil {
			return Pokemon{}, err
		}
		return pokemonRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Pokemon{}, fmt.Errorf("failed to fetch pokemon '%s'", pokemonName)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonRes := Pokemon{}
	if err := json.Unmarshal(data, &pokemonRes); err != nil {
		return Pokemon{}, err
	}

	c.pokeCache.Add(url, data)

	return pokemonRes, nil
}
