package main

import (
	"fmt"
)

func main() {
	// Define a map of countries and their capitals
	countries := map[string]string{
		"France":      "Paris",
		"Germany":     "Berlin",
		"Italy":       "Rome",
		"Spain":       "Madrid",
		"Japan":       "Tokyo",
		"Canada":      "Ottawa",
		"Argentina":   "Buenos Aires",
		"South Korea": "Seoul",
	}

	// Loop through the map and print in desired format
	for country, capital := range countries {
		fmt.Printf("Capital of %s is %s\n", country, capital)
	}
}
