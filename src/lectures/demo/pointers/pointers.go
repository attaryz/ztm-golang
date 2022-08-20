package main

import "fmt"

type Counter struct {
	hits int
}

func increment(counter *Counter) {
	counter.hits += 1

	fmt.Println("Counter", counter)
}

func replace(old *string, new string, counter *Counter) {
	*old = new
	increment(counter)
}

func main() {

	counter := Counter{}

	hello := "hello"
	world := "world"

	fmt.Println("hello", hello)
	fmt.Println("world", world)

	replace(&hello, "hi", &counter)
	fmt.Println(hello, world)

	phrase := []string{hello, world}

	fmt.Println(phrase)

	replace(&phrase[1], "yes", &counter)
	fmt.Println(phrase)
}
