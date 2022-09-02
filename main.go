package main

// format package
import "fmt"

func main() {
	var conferenceName = "Go Conference"
	const conferenceTickets = 50
	var remainingTickets = conferenceTickets

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Println("We have total", conferenceTickets, "tickets", remainingTickets, "are still available")
	fmt.Println("Get your tickets here to attend")
}
