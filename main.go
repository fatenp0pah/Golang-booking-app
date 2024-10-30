package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50
	bookings := []string{} // size of array
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T, conferenceName is %T\n", conferenceTickets, remainingTickets, conferenceTickets)

	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have total of %v ticket and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint //uint positive

		// ask user their name
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) // scan user input
		// fmt.Print(&remainingTickets)
		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName) // scan user input

		fmt.Println("Enter your email: ")
		fmt.Scan(&email) // scan user input

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets) // scan user input

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v tickets remaining, so you can't book %v tickets\n", remainingTickets, userTickets)
			continue
		}
		// userTickets = 2
		// fmt.Printf("User %v booked %v tickets.\n", userName, userTickets)

		remainingTickets = remainingTickets - userTickets
		// bookings[0] = firstName + " " + lastName
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("The whole slice: %v\n", bookings)
		fmt.Printf("The first value: %v\n", bookings[0])
		fmt.Printf("Slice type: %T\n", bookings)
		fmt.Printf("Slice length: %v\n", len(bookings))

		fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			// var firstName = names[0]
			firstNames = append(firstNames, names[0])
		}

		fmt.Printf("The first names of bookings: %v\n", firstNames)

		noTicketsRemaining := remainingTickets == 0
		if noTicketsRemaining {
			// end program
			fmt.Println("Our conference is booked out. Comeback next year.")
			break
		}

	}

}
