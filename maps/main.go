package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"white": "#ffffff",
	}

	printColors(colors)
}

func printColors(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("Hex code for color %v is %v\n", color, hex)
	}
}
