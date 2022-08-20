package main

import "fmt"

type Passenger struct {
	Name         string
	TicketNumber int
	Boarded      bool
}

type Bus struct {
	FrontSeat Passenger
}

func main() {

	casey := Passenger{"Casey", 1, false}
	fmt.Println(casey)

	var (
		bill = Passenger{Name: "Bill", TicketNumber: 2}
		ella = Passenger{Name: "Ella", TicketNumber: 3}
	)
	fmt.Println(bill, ella)

	bill.Boarded = true
	if bill.Boarded {
		fmt.Println("Bill is on the bus")
	}
	if casey.Boarded {
		fmt.Println("Casey is on the bus")
	}

}
