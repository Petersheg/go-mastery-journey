package main

import (
	"fmt"
	"log"
	"github.com/Petersheg/go-mastery-journey/day06-module-conclusion/greetings"
)

// main is the entry point for the application.
func main() {
	log.SetPrefix("Greetings06 ")
	log.SetFlags(0)
	
	names := []string{"Ajayi Oluokun", "Ajagbe Bintu", "Onakoyo Adeleke"}

	message, error := greetings06.Hellos(names);

	if(error != nil){
		log.Fatal(error)
	}

	fmt.Println(message)
}
