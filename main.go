package main

// format package
import "fmt"

func main() {
	// := is sugar syntax for `var conferenceName string = "conference"`, this doesn't work for const
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	remainingTickets := conferenceTickets

	// %T: type
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	// %v: value
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")

	var userName string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	// Scan to ask for prompt user input
	// pointer & is a variable that points to the memory address of another variable
	// Scan can assign the user's value to the userName variable cuz it has a pointer to its memory address
	fmt.Scan(&userName)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets

	fmt.Printf("User %v booked %v tickets. %v tickets left.\n", userName, userTickets, remainingTickets)

}
