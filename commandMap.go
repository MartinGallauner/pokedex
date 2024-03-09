package main

import "fmt"

func commandMapf() error {
	response := getLocations()
	for _, val := range response.Results {
		fmt.Println(val.Name)
	}
	return nil
}

func commandMapb() error {
	response := getLocations()
	for _, val := range response.Results {
		fmt.Println(val.Name)
	}
	return nil
}
