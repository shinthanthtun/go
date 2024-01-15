package helper

import (
	"strings"
)

func ValidateUserInput(firstName string, lastName string, email string, userTickets int, remainingTicket int) (bool, bool, bool) {
	isValidName := len(firstName) > 2 && len(lastName) > 2
	isValidEmail := strings.Contains(email, "@")
	isValidTickets := userTickets > 0 && userTickets <= remainingTicket

	return isValidName, isValidEmail, isValidTickets
}
