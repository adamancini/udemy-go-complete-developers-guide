package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#0000ff",
		"green": "#00ff00",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	// colors["white"] = "#ffffff"

	// fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex for ", color, "is ", hex)
	}
}
