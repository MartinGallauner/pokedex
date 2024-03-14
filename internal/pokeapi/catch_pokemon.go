package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) CatchPokemon(parameter string) (PokemonResponse, error) {
	url := baseURL + "/pokemon/" + parameter

	if val, ok := c.cache.Get(url); ok {
		response := PokemonResponse{}
		err := json.Unmarshal(val, &response)
		if err != nil {
			return PokemonResponse{}, err
		}
		return response, nil
	}

	req, error := http.NewRequest("GET", url, nil)
	if error != nil {
		return PokemonResponse{}, error
	}

	resp, error := c.httpClient.Do(req)
	if error != nil {
		return PokemonResponse{}, error
	}
	defer resp.Body.Close()

	dat, error := io.ReadAll(resp.Body)
	if error != nil {
		return PokemonResponse{}, error
	}

	response := PokemonResponse{}
	error = json.Unmarshal(dat, &response)
	if error != nil {
		return PokemonResponse{}, error
	}
	c.cache.Add(url, dat)
	return response, nil
}
