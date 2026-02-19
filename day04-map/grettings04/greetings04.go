package grettings04

import (
	"math/rand"
	"errors"
	"fmt"
)

// Hello returns a greeting message for the named person.
func Hello(name string) (string, error) {
	if(name == "") {
		return "", errors.New("Name cannot be empty")
	}
	message := fmt.Sprintf(rabdomGreeting(), name)

	return message, nil
}


// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string] string, error) {

	// A map to associate names with messages.
	messages := make(map[string] string)

	// Loop through the received slice of names, calling
    // the Hello function to get a message for each name.
	for _, name := range names {
		message, error := Hello(name)

		if(error != nil) {
			return nil, error
		}

		// In the map, associate the retrieved message with
        // the name.
		messages[name] = message
	}

	return  messages, nil

}

// rabdomGreeting returns a random greeting message.
func rabdomGreeting() string {

	messages := []string {
		"Hi there, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	i := rand.Intn(len(messages))
	return messages[i]
}