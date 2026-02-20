package greetings05

import (
	"fmt"
	"errors"
	"math/rand"
)

func Hello(name string) (string, error) {

	if(name == "") {
		return "", errors.New("Name can not be empty")
	}

	message := fmt.Sprintf(randomGreeting(), name)

	return  message, nil
}

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

func randomGreeting() string {

	messages := []string {
		"Hello, %v, you are welcome",
		"Glad to have you here, %v.",
		"%v, Peace be unto to you!",
	}

	message := messages[rand.Intn(len(messages))]

	return message
}
