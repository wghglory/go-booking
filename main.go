package main

// format package
import (
	"fmt"
	"strings"
)

// := is sugar syntax for `var conferenceName string = "conference"`, this doesn't work for const
var conferenceName = "Go Conference"

const conferenceTickets uint = 50

var remainingTickets = conferenceTickets

// var bookings = [50]string{"Derek", "Nana", "Peter"}
// var bookings = [50]string{}
// var bookings []string // slice
// var bookings = []string{} // slice
var bookings = []string{} // slice

func main() {

	// %T: type
	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	greetUsers()

	for {
		firstName, lastName, userTickets := getUserInputs()

		if len(firstName) <= 2 || len(lastName) <= 2 {
			fmt.Println("firstName and lastName length must be more than 2 characters. Please try again.")
			continue
		}

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, you can't book %v tickets.\n", remainingTickets, userTickets)
			continue
		}

		remainingTickets, bookings = bookTicket(userTickets, firstName, lastName)

		firstNames := getFirstNames()
		fmt.Printf("FirstNames of all our bookings: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out.")
			break
		}
	}

}

func bookTicket(userTickets uint, firstName string, lastName string) (uint, []string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("User %v booked %v tickets. %v tickets left.\n", firstName, userTickets, remainingTickets)
	return remainingTickets, bookings
}

func getUserInputs() (string, string, uint) {
	var firstName string
	var lastName string
	var userTickets uint

	fmt.Println("Enter your firstName: ")
	// Scan to ask for prompt user input
	// pointer & is a variable that points to the memory address of another variable
	// Scan can assign the user's value to the firstName variable cuz it has a pointer to its memory address
	// validation
	fmt.Scan(&firstName)

	fmt.Println("Enter your lastName: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)
	return firstName, lastName, userTickets
}

func getFirstNames() []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func greetUsers() {
	// %v: value
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have total %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")
}
