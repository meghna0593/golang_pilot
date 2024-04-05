package main

import "fmt"

func main() {

	// First way
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}

	// Second way
	// var colors map[string]string

	// Third way
	// colors := make(map[string]string)

	colors["white"] = "#ffffff"

	// Delete keys and values
	// delete(colors, "white")
	printMap(colors)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
