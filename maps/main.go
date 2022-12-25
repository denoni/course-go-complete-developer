package main

import "fmt"

func main() {
	colors := make(map[string]string)

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#4bf745",
	// }

	colors["white"] = "#ffffff"
	colors["black"] = "#000000"

	delete(colors, "white")

	printMap(colors)
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println("Key:", key)
		fmt.Println("Value:", value)
	}
}
