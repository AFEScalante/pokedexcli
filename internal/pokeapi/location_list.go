package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	cachedData, cacheExists := c.pokeCache.Get(url)
	if cacheExists {
		locationsRes := RespShallowLocations{}
		if err := json.Unmarshal(cachedData, &locationsRes); err != nil {
			return RespShallowLocations{}, err
		}
		return locationsRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}
	
	locationsRes := RespShallowLocations{}
	if err := json.Unmarshal(data, &locationsRes); err != nil {
		return RespShallowLocations{}, err
	}

	c.pokeCache.Add(url, data)

	return locationsRes, nil
}
