package main

import (
	"fmt"
	"log"
	"github.com/Petersheg/go-mastery-journey/day03-erro-handling/greetings03"
)

func main() {
	// Set properties of the predefined Logger, including
    // the log entry prefix and a flag to disable printing
    // the time, source file, and line number.
	log.SetPrefix("greetings03: ")
	log.SetFlags(0)

	message, err := greetings03.Hello("Agbeloba Jacob")
	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message
    // to the console.
	fmt.Println(message)
}
