package main

import "fmt"

func commandMap() error {
	locations := getLocations()
	for _, val := range locations.Results {
		fmt.Println(val.Name)
	}
	return nil
}
