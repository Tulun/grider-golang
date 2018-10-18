package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":  "#FF0000",
		"blue": "#0000FF",
	}

	colors["white"] = "#ffffff"
	fmt.Println(colors)
	fmt.Println(colors["white"])
	printMap(colors)
	delete(colors, "white")
	fmt.Println(colors)
	var blerp map[string]string
	fmt.Println(blerp)

	bloop := make(map[string]string)
	fmt.Println(bloop)

}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
