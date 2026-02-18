package greetings03

import (
	"math/rand"
	"errors"
	"fmt"
)

// Hello returns a greeting message for the named person.
func Hello (name string) (string, error) {
	if name == "" {
		return "", errors.New("Empty name")
	}
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func randomGreeting() string {
	//A slice of messages to use as a template for the greeting.
	messages := []string {
		"Hi there, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	// Return a randomly selected greeting message by specifying
	// a random index for the slice of messages.
	i := rand.Intn(len(messages))
	return messages[i]
}