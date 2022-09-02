package main

// format package
import "fmt"

func main() {
	// := is sugar syntax for `var conferenceName string = "conference"`, this doesn't work for const
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	remainingTickets := conferenceTickets

	// %T: type
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	// %v: value
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var userName string
	var userTickets int

	userName = "Derek"
	userTickets = 2
	fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

}
