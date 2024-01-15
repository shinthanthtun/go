package main

import (
	"fmt"
)

var conferenceName string = "Music Conference"

const conferenceTicket int = 50

var remainingTicket uint = 50
var bookings []string

func main() {

	for remainingTicket > 0 && len(bookings) < 50 {
		greetUser()

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTickets := validateUserInput(firstName, lastName, email, userTickets)

		if isValidName && isValidEmail && isValidTickets {
			bookTickets(userTickets, firstName, lastName, email)
			//Print first name function
			firstNames := printFirstName(bookings)
			fmt.Printf("These are all our bookings : %v\n", firstNames)

			if remainingTicket == 0 {
				fmt.Println("Our conference is booked out. Come back next year!\n")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name you entered is too short!")
			}
			if !isValidEmail {
				fmt.Println("Email you entered lack @ sign!")
			}
			if !isValidTickets {
				fmt.Println("Number of tickets you entered is invalid")
			}
		}

	}
}
