package main

import "fmt"

func price() int {
	return 1
}

const (
	Economy    = 0
	Business   = 1
	FirstClass = 2
)

func main() {

	switch p := price(); {
	case p < 2:
		fmt.Println("cheap item")
	case p < 10:
		fmt.Println("normal item")
	default:
		fmt.Println("expensive item")
	}

	ticket := Economy
	switch ticket {
	case Economy:
		fmt.Println("Economy")
	case Business:
		fmt.Println("Business")
	case FirstClass:
		fmt.Println("FirstClass")
	default:
		fmt.Println("Invalid ticket")
	}
}
