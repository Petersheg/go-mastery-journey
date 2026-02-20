package main

import (
	"fmt"
	"log"
	"github.com/Petersheg/go-mastery-journey/day05-test/greetings"
)

func main() {
	log.SetPrefix("Greetings05 ")
	log.SetFlags(0)
	
	names := []string{"Ajayi Oluokun", "Ajagbe Bintu", "Onakoyo Adeleke"}

	message, error := greetings05.Hellos(names);

	if(error != nil){
		log.Fatal(error)
	}

	fmt.Println(message)
}
