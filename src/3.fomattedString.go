package main

import "fmt"

func main() {

	//string formatting ways
	//%v = interpolate the default representation
	var myName string
	myName = "Shin Thant Htun"
	fmt.Printf("I am %v \n", myName)

	//%s = interpolate a string
	//We can just use %v instead of %s and %d
	fmt.Printf("I %s Khin Sandi Ko \n", "Love You")

	//%d = interpolate an integer in decimal form
	fmt.Printf("Now I am %d years old \n", 18)

	//%f = interpolate a decimal
	//If you want a 2 decimal value you can just use %.2f
	fmt.Printf("I am %.2f years old \n", 18.5688)
	// the output is just 18.57 , you round the last number

	//Little Assignment

	const fullName = "Shin Thant Htun"
	const openRate = 100.65

	msg := fmt.Sprintf("Hi %s , your open rate is %.1f", fullName, openRate)
	fmt.Print(msg)

}
