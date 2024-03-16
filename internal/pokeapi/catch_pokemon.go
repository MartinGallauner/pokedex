package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(parameter string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + parameter

	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		return Pokemon{}, error
	}

	resp, error := c.httpClient.Do(req)
	if error != nil {
		return Pokemon{}, error
	}
	defer resp.Body.Close()

	dat, error := io.ReadAll(resp.Body)
	if error != nil {
		return Pokemon{}, error
	}

	response := Pokemon{}
	error = json.Unmarshal(dat, &response)
	if error != nil {
		return Pokemon{}, error
	}
	return response, nil
}
