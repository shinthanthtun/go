package main

import (
	"booking-app/helper"
	"fmt"
	"strings"
)

var conferenceName string = "Music Conference"

const conferenceTicket int = 50

var remainingTicket int = 50
var Bookings []string

func main() {

	for remainingTicket > 0 && len(Bookings) < 50 {
		GreetUser(conferenceName, conferenceTicket, remainingTicket)

		firstName, lastName, email, userTickets := GetUserInput()

		isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTicket)

		if isValidName && isValidEmail && isValidTickets {
			BookTickets(userTickets, firstName, lastName, email)
			//Print first name function
			firstNames := PrintFirstName(Bookings)
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
			break
		}

	}
}

func GreetUser(conferenceName string, conferenceTicket int, remainingTicket int) {
	fmt.Printf("\nWelcome to %v booking application!!!\n", conferenceName)
	fmt.Printf("We have total of %v tickens and %v are still available.\n", conferenceTicket, remainingTicket)
	fmt.Println("Get your tickets here to attend")
}

func GetUserInput() (string, string, string, int) {
	var firstName string
	var lastName string
	var email string
	var userTickets int

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
func BookTickets(userTickets int, firstName string, lastName string, email string) {
	remainingTicket = remainingTicket - userTickets
	Bookings = append(Bookings, firstName+" "+lastName)

	fmt.Printf("Thank you %v %v for booking.\n", firstName, lastName)
	fmt.Printf("You will reveice a comfirmation email at your mail address %v\n", email)
}
func PrintFirstName(bookings []string) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		var names = strings.Fields(booking)
		firstNames = append(firstNames, names[0]+", ")
	}
	return firstNames
}
