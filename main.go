package main

import(
  "fmt"
)

func main(){
  var conferenceName = "Music Conference"
  const conferenceTicket = 50
  var remainingTicket = 50


  fmt.Printf("Welcome to %v booking application!!!\n" , conferenceName)
  fmt.Printf("We have total of %v tickens and %v are still available.\n", conferenceName , remainingTicket)
  fmt.Println("Get your tickets here to attend")
  
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

  fmt.Printf("Thank you %v %v for booking.\n" , firstName, lastName)
  fmt.Printf("You will reveice a comfirmation email at your mail address %v\n" , email)
  
}

