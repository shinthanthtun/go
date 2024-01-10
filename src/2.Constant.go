package main

import "fmt"

func main() {
	//In constant we can't use short declaration syntax like (:=)
	//We can't overwrite constant variables

	// const myName = "Shin Thant Htun"
	// fmt.Println(myName)

	const firstName = "Shin Thant"
	const lastName = "Htun"
	const fullName = firstName + " " + lastName
	fmt.Print(fullName)
}
