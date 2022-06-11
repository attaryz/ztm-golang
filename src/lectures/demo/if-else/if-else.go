package main

import "fmt"

func average(a, b, c int) float32 {
	// Convert the sum of the scores into a float32
	return float32(a+b+c) / 3
}

func main() {

	quiz1, quiz2, quiz3 := 9, 7, 8

	if quiz1 > quiz2 {
		fmt.Println("quiz1 scored higher than quiz1")

	} else if quiz1 < quiz2 {
		fmt.Println("quiz1 scored lower than quiz1")
	} else {
		fmt.Println("quiz1 scored same as quiz2")
	}

	if average(quiz1, quiz2, quiz3) > 7 {
		fmt.Println(("acceptable score"))
	} else {
		fmt.Println(("unacceptable score"))
	}
}
