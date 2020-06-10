package main

import "fmt"

func main() {
	colours := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
		"blue":  "#0000ff",
		"white": "#ffffff",
		"black": "#000000",
	}

	printColours(colours)

	fmt.Println("ADDING ORANGE")
	colours["orange"] = "#ffa500"

	printColours(colours)
}

func printColours(c map[string]string) {
	for colour, hex := range c {
		fmt.Printf("The hexadecimal for colour %v is %v\n", colour, hex)
	}
}
