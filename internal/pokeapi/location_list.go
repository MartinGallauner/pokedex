package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationsResponse, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		response := LocationsResponse{}
		err := json.Unmarshal(val, &response)
		if err != nil {
			return LocationsResponse{}, err
		}
		return response, nil
	}

	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		return LocationsResponse{}, error
	}

	resp, error := c.httpClient.Do(req)
	if error != nil {
		return LocationsResponse{}, error
	}
	defer resp.Body.Close()

	dat, error := io.ReadAll(resp.Body)
	if error != nil {
		return LocationsResponse{}, error
	}

	response := LocationsResponse{}
	error = json.Unmarshal(dat, &response)
	if error != nil {
		return LocationsResponse{}, error
	}
	c.cache.Add(url, dat)
	return response, nil
}
