package greetings

import "fmt"

// at func Hello (name string) string .
// name = parameter variable
// string = parameter variable type
// another string is return type

func Hello(name string) string {

	// := is a operator used as shortcut for declaring and initializing a variable in one line
	// if we don't use := we will type like
	// var message string
	// message = fmt.Sprintf("Hello, %v . Welcome")

	// fmt.Sprintf is (save formatted string and return the string to us)
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
