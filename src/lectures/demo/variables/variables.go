package main

import "fmt"

func main() {
	var myName = "Abdullah"
	fmt.Println("my name is", myName)

	var name string = "ali"
	fmt.Println("name =", name)

	username := "attaryz"
	fmt.Println("username =", username)

	var total int
	fmt.Println("total =", total)

	movieName, movieYear := "Avengers: End Game", 2019 // multiple assignment
	fmt.Println("movie name:", movieName, ", movie year:", movieYear)

	var (
		playerName  = "Paulo Dybala"
		teamName    = "Juventus"
		shirtNumber = 10
	)
	fmt.Println("player name:", playerName, ", team name:", teamName, ", shirt number:", shirtNumber)

}
