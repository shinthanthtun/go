package main

import (
	"booking-app/helper"
	"fmt"
	"sync"
	"time"
)

var conferenceName string = "Music Conference"

const conferenceTicket int = 50

var remainingTicket int = 50
var Bookings = make([]UserData, 0)

type UserData struct {
	firstName   string
	lastName    string
	email       string
	noOfTickets int
}

var wg = sync.WaitGroup{}

func main() {

	GreetUser(conferenceName, conferenceTicket, remainingTicket)

	firstName, lastName, email, userTickets := GetUserInput()

	isValidName, isValidEmail, isValidTickets := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTicket)

	if isValidName && isValidEmail && isValidTickets {
		BookTickets(userTickets, firstName, lastName, email)

		wg.Add(1)
		go sendTickets(userTickets, firstName, lastName, email)
		//Print first name function
		firstNames := PrintFirstName()
		fmt.Printf("These are all our bookings : %v\n", firstNames)

		if remainingTicket == 0 {
			fmt.Println("Our conference is booked out. Come back next year!")
			//break
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
		//break

	}
	wg.Wait()
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

	//Create a map for a user data

	var userDate = UserData{
		firstName:   firstName,
		lastName:    lastName,
		email:       email,
		noOfTickets: userTickets,
	}

	Bookings = append(Bookings, userDate)
	fmt.Printf("Thank you %v %v for booking.\n", firstName, lastName)
	fmt.Printf("You will reveice a comfirmation email at your mail address %v\n", email)
}
func PrintFirstName() []string {
	firstNames := []string{}
	for _, booking := range Bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func sendTickets(userTickets int, firstName string, lastName string, email string) {
	time.Sleep(5 * time.Second)
	tickets := fmt.Sprintf("%v ticekts for %v %v", userTickets, firstName, lastName)
	fmt.Println("********************")
	fmt.Printf("Sending Tickets:\n%v\nTo email address : %v\n", tickets, email)
	fmt.Println("********************")
	wg.Done()
}
