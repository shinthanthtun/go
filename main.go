package main

import (
	"fmt"
	"strings"
)

func main() {
	var conferenceName string = "Music Conference"
	const conferenceTicket int = 50
	var remainingTicket uint = 50
	var bookings []string

	fmt.Printf("Welcome to %v booking application!!!\n", conferenceName)
	fmt.Printf("We have total of %v tickens and %v are still available.\n", conferenceName, remainingTicket)
	fmt.Println("Get your tickets here to attend")

	for remainingTicket > 0 && len(bookings) < 50 {
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

		isValidName := len(firstName) > 2 && len(lastName) > 2
		isValidEmail := strings.Contains(email, "@")
		isValidTickets := userTickets > 0 && userTickets <= remainingTicket

		if isValidName && isValidEmail && isValidTickets {
			remainingTicket = remainingTicket - userTickets
			bookings = append(bookings, firstName+" "+lastName)

			firstNames := []string{}
			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0])
			}

			fmt.Printf("Thank you %v %v for booking.\n", firstName, lastName)
			fmt.Printf("You will reveice a comfirmation email at your mail address %v\n", email)
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
