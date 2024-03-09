package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func getLocations() LocationResponse {
	res, err := http.Get("http://pokeapi.co/api/v2/location")
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	response := LocationResponse{}
	err = json.Unmarshal(body, &response)

	if err != nil {
		fmt.Println(err)
	}
	return response
}

type LocationResponse struct {
	Count    int       `json:"count"`
	Next     string    `json:"next"`
	Previous string    `json:"previous"`
	Results  []Results `json:"results"`
}
type Results struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
