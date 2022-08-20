package main

import "fmt"

type Room struct {
	name    string
	cleaned bool
}

func checkCleanliness(rooms [4]Room) {
	for i := 0; i < len(rooms); i++ {
		room := rooms[i]
		if room.cleaned {
			fmt.Println(room.name, "is clean")
		} else {
			fmt.Println(room.name, "is not clean")
		}
	}
}
func main() {
	rooms := [...]Room{
		{name: "Office"},
		{name: "Room"},
		{name: "Warehouse"},
		{name: "Store"},
	}
	checkCleanliness(rooms)
	fmt.Println("cleaning...")
	rooms[2].cleaned = true
	rooms[3].cleaned = true
	checkCleanliness(rooms)
}
