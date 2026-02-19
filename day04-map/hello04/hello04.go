package main

import (
	"fmt"
	"log"
	"github.com/Petersheg/go-mastery-journey/day04-map/grettings04"
)

func main() {
	//slice of name
	names := []string{"Peter oluwasegun", "Abiola Adele", "Joy Akanji"};


	greeting, error := grettings04.Hellos(names)

	log.SetPrefix("Greetings04: ")
	log.SetFlags(0)

	if(error != nil) {
		log.Fatal(error)
	}

	fmt.Println(greeting)
}