//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

import "fmt"

func greetPerson(personName string) string {
	return fmt.Sprintf("Hello %s", personName)
}
func functionThatReturnsAnyMessage() string {
	return "Hello"
}

func addThreeNumbers(x int, y int, z int) int {
	return x + y + z
}

func functionThatReturnsAnyNumber() int {
	return 1
}
func functionThatReturnsAnyTwoNumbers() (int, int) {
	return 1, 2
}

func main() {
	fmt.Println(greetPerson("Abdullah"))
	fmt.Println(functionThatReturnsAnyMessage())
	fmt.Println(addThreeNumbers(1, 2, 3))
	fmt.Println(functionThatReturnsAnyNumber())
	fmt.Println(functionThatReturnsAnyTwoNumbers())

	addNumber := addThreeNumbers(functionThatReturnsAnyNumber(), 10, 20)
	fmt.Println(addNumber)
}
