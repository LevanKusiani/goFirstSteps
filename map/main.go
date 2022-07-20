package main

import "fmt"

func main() {
	// v1 of creating a map
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
	}

	// v2 of creating a map
	// var colors map[string]string

	// v3 of creating a map
	// colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["grey"] = "#ffbbcc"

	delete(colors, "grey")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println(color + ": " + hex)
	}
}
