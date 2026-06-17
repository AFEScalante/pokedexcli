package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(area_name string) (Location, error) {
	url := baseURL + "/location-area/" + area_name

	cachedData, cacheExists := c.pokeCache.Get(url)
	if cacheExists {
		locationRes := Location{}
		if err := json.Unmarshal(cachedData, &locationRes); err != nil {
			return Location{}, err
		}
		return locationRes, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return Location{}, fmt.Errorf("failed to fetch location for area '%s'", area_name)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationRes := Location{}
	if err := json.Unmarshal(data, &locationRes); err != nil {
		return Location{}, err
	}

	c.pokeCache.Add(url, data)

	return locationRes, nil
}
