package main

import (
	"errors"
	"log"
)

func validateInt(number int) (int, error) {

	if(number <= 0) {
		return  0, errors.New("Provide a number grether than Zero")
	}

	return number, nil
}

/*
	Go has only one looping construct, the for loop.
	The basic for loop has three components separated by semicolons:

	the init statement: executed before the first iteration
	the condition expression: evaluated before every iteration
	the post statement: executed at the end of every iteration
	The init statement will often be a short variable declaration, and the variables
	declared there are visible only in the scope of the for statement.

	The loop will stop iterating once the boolean condition evaluates to false.

	Note: Unlike other languages like C, Java, or JavaScript there are no parentheses surrounding
	the three components of the for statement and the braces { } are always required.
*/
func SumUp(number int) int {
	sum := 0
	_, err := validateInt(number)

	if(err != nil) {
		log.Fatal(err)
	}

	for i := 0; i <= number; i++{
		sum += i
	}

	return  sum
}

// The init and post statements are optional.
func SumUpWithoutInit(number int) int {
	sum := 1
	_, err := validateInt(number)

	if(err != nil){
		log.SetPrefix("For: ")
		log.SetFlags(1)

		log.Fatal(err)
	}

	for ; sum < 200; {
		sum += sum
	}

	return sum
}

/*
	For is Go's "while"
	At that point you can drop the semicolons: C's while is spelled for in Go.	
*/
func ForAsWhile() int {
 	sum := 1

	// This is while loop
	for sum < 100 {
		sum += sum
	}

	return sum
}

/*
	Forever
	If you omit the loop condition it loops forever, so an infinite loop is compactly expressed.
*/
func Forever(){
	// for{}
}