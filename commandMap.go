package main

import "fmt"

func commandMap() error {
	locations := getLocations()
	fmt.Println(locations)

	return nil
}
