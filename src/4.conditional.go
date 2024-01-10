package main

import "fmt"

func main() {

	//if statements in Go don't use parentheses arount the condition
	//And watch the else syntax down below , syntax is important
	messageLen := 10
	maxMessageLen := 20
	fmt.Println("Trying to send a message of length : ", messageLen)

	if messageLen <= maxMessageLen {
		fmt.Println("Message sent")
	} else {
		fmt.Println("Message not sent")
	}

	/*
		Why would  I use this?
		This is just some syntactic sugar that Go offer to shorten up code in some
		cases. For example , insted of writing

		length := getLength(email)
		if length < 1{
			fmt.Println("Email is invalid")
		}

		we can do

		if length := getLength(email); length < 1 {
			fmt.Println("Email is invalid")
		}
	*/

}
