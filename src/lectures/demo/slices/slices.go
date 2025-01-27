package main

import "fmt"

func printSlice(title string, slice []string) {
	fmt.Println()
	fmt.Println("---", title, "---")
	for i := 0; i < len(slice); i++ {
		element := slice[i]
		fmt.Println(element)
	}
}

func main() {
	route := []string{"Bank", "Barber", "Grocery", "Hardware"}
	printSlice("Route", route)
	route = append(route, "Home")
	printSlice("2nd Route", route)
	fmt.Println()
	fmt.Println(route[0], "visited")
	fmt.Println(route[1], "visited")

	route = route[2:]
	printSlice("remaining locations", route)
}
