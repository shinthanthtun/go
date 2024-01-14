package main

import (
	"fmt"
)

func main() {
	var conferenceName string = "Music Conference"
	const conferenceTicket int = 50
	var remainingTicket uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application!!!\n", conferenceName)
	fmt.Printf("We have total of %v tickens and %v are still available.\n", conferenceName, remainingTicket)
	fmt.Println("Get your tickets here to attend")

	for {
		var firstName string
		var lastName string
		var email string
		var userTickets uint

		//ask user for their tickets
		fmt.Print("What's your first name : ")
		fmt.Scan(&firstName)

		fmt.Print("What's your last name : ")
		fmt.Scan(&lastName)

		fmt.Print("What's your email : ")
		fmt.Scan(&email)

		fmt.Print("How many ticket do you wanna book : ")
		fmt.Scan(&userTickets)

		remainingTicket = remainingTicket - userTickets
		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("Thank you %v %v for booking.\n", firstName, lastName)
		fmt.Printf("You will reveice a comfirmation email at your mail address %v\n", email)
		fmt.Printf("These are all our bookings : %v\n", bookings)
	}
}
