package main

import (
	"fmt"
	"github.com/Petersheg/go-mastery-journey/day02-create-module/greetings"
)

func main() {
	message := greetings.Hello("Oluwasegun")
	fmt.Println(message)
}