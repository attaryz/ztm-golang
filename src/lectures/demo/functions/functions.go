package main

import "fmt"

func double(x int) int {
	return x * 2
}

func add(lhs, rhs int) int {
	return lhs + rhs
}

func greet() {
	fmt.Println(("greet function!!"))
}

func main() {

	greet()
	dozen := double(4)
	fmt.Println("dozin", dozen)

	bakerDozin := add(dozen, 1)
	fmt.Println("bakerDozin", bakerDozin)

}
