package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocationData(location string) (RespExactLocation, error) {
	url := baseURL + "/location-area/" + location
	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespExactLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespExactLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespExactLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespExactLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespExactLocation{}, err
	}

	exactLocationResp := RespExactLocation{}
	err = json.Unmarshal(dat, &exactLocationResp)
	if err != nil {
		return RespExactLocation{}, err
	}

	c.cache.Add(url, dat)
	return exactLocationResp, nil
}
