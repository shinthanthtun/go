package main

import (
	"fmt"
	"strings"
)

func greetUser() {
	fmt.Printf("\nWelcome to %v booking application!!!\n", conferenceName)
	fmt.Printf("We have total of %v tickens and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")
}

func printFirstName(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0])
	}
	return firstNames
}

func validateUserInput(firstName string, lastName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTicket

	return isValidName, isValidEmail, isValidTickets
}

func getUserInput() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets
}

func bookTickets(userTickets uint, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTickets
	bookings = append(bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking.\n", firstName, lastName)
	fmt.Printf("You will reveice a comfirmation email at your mail address %v\n", email)

}
