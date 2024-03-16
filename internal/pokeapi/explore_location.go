package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(pageURL *string, parameter string) (LocationResponse, error) {
	url := baseURL + "/location-area/" + parameter

	if val, ok := c.cache.Get(url); ok {
		response := LocationResponse{}
		err := json.Unmarshal(val, &response)
		if err != nil {
			return LocationResponse{}, err
		}
		return response, nil
	}

	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		return LocationResponse{}, error
	}

	resp, error := c.httpClient.Do(req)
	if error != nil {
		return LocationResponse{}, error
	}
	defer resp.Body.Close()

	dat, error := io.ReadAll(resp.Body)
	if error != nil {
		return LocationResponse{}, error
	}

	response := LocationResponse{}
	error = json.Unmarshal(dat, &response)
	if error != nil {
		return LocationResponse{}, error
	}
	c.cache.Add(url, dat)
	return response, nil
}
