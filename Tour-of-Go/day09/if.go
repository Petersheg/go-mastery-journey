package main

import (
	"fmt"
	"math/rand"
)

/*
	Go's if statements are like its for loops; the expression need
	not be surrounded by parentheses ( ) but the braces { } are required.
*/ 
func IfFunc (num int){
	if num == 0 {
		fmt.Println("Neutral")
	}

	fmt.Printf("Your input is %v\n", num)
}

/*
	If with a short statement
	Like for, the if statement can start with a short statement 
	to execute before the condition.

	Variables declared by the statement are only in scope until the end of the if.
*/
func IfWithStatement () {

	if num := 0; num < 0 {
		fmt.Println("Negative")
	}

    fmt.Println("Positive number")
}

/*
	Variables declared inside an if short statement are also available 
	inside any of the else blocks.
*/ 
func IfAndElse (guess int) {
	rand := rand.Intn(5);

	if num := guess; num == rand {
		fmt.Printf("You gues right %v\n", rand)
	}else {
		fmt.Printf("Your guess is wrong, the right number is %v\n", rand)
	}

	fmt.Println("Game end.")
}
