package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (LocationResponse, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
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
	return response, nil
}
