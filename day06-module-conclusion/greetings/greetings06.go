package greetings06

import (
	"fmt"
	"errors"
	"math/rand"
)

// Hello returns a greeting for the named person. If the name is empty, an error is returned.
func Hello(name string) (string, error) {

	if(name == "") {
		return "", errors.New("Name can not be empty")
	}

	message := fmt.Sprintf(randomGreeting(), name)

	return  message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message. If a name is empty, an error is returned.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string);

	for _, name := range names {
		message, error := Hello(name)

		if(error != nil) {
			return nil, error
		}

		messages[name] = message
	}

	return messages, nil
}

// randomGreeting returns a random greeting message.
func randomGreeting() string {

	messages := []string {
		"Hello, %v, you are welcome",
		"Glad to have you here, %v.",
		"%v, Peace be unto to you!",
	}

	message := messages[rand.Intn(len(messages))]

	return message
}
