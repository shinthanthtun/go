package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Hein Htet Zaw")
	fmt.Println(message)
}
