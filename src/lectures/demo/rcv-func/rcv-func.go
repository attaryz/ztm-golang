package main

import "fmt"

type Space struct {
	occupied bool
}

type ParkingLot struct {
	spaces []Space
}

func occupySpace(lot *ParkingLot, spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}

func (lot *ParkingLot) occupySpaceReceiver(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = true
}
func (lot *ParkingLot) vacateSpaceReceiver(spaceNum int) {
	lot.spaces[spaceNum-1].occupied = false
}

func main() {

	lot := ParkingLot{spaces: make([]Space, 5)}

	fmt.Println("Initial Space: ", lot)

	lot.occupySpaceReceiver(1)
	lot.occupySpaceReceiver(5)
	occupySpace(&lot, 2)
	fmt.Println("After Occupy: ", lot)
	lot.vacateSpaceReceiver(1)
	fmt.Println("After vacate: ", lot)

}
