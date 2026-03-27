package main

import (
	"fmt"
	"math"
	"time"
)

/* Errors
Go programs express error state with error values.

The error type is a built-in interface similar to fmt.Stringer:

type error interface {
	Error() string
}
(As with fmt.Stringer, the fmt package looks for the error interface when printing values.)

Functions often return an error value, and calling code should handle errors by testing whether the error equals nil.

i, err := strconv.Atoi("42")
if err != nil {
	fmt.Printf("couldn't convert number: %v\n", err)
	return
}
fmt.Println("Converted integer:", i)
A nil error denotes success; a non-nil error denotes failure.
*/ 
type MyError struct {
	When time.Time
	What string
}

func (e MyError) Error() string{
	return fmt.Sprintf("at %v, %s", e.What, e.When)
}

func run() error {
	return &MyError{
		time.Now(),
		"It did not work",
	}
}

func Testrun() {
	if err := run(); err != nil {
		fmt.Println("Error:", err)
	}
}


/* Exercise: Errors
	Copy your Sqrt function from the earlier exercise and modify it to return an error value.

	Sqrt should return a non-nil error value when given a negative number, as it doesn't support complex numbers.

	Create a new type

	type ErrNegativeSqrt float64
	and make it an error by giving it a

	func (e ErrNegativeSqrt) Error() string
	method such that ErrNegativeSqrt(-2).Error() returns "cannot Sqrt negative number: -2".

	Note: A call to fmt.Sprint(e) inside the Error method will send the program into an infinite loop. You can avoid this by converting e first: fmt.Sprint(float64(e)). Why?

	Change your Sqrt function to return an ErrNegativeSqrt value when given a negative number.

*/ 

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}
	
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}

	z := 1.0
	const epsilon = 0.0000001

	for range 10 {
		prev_z := z

		z -= (z*z - x) / (2*z)
		if(math.Abs(prev_z - z) < epsilon){
			break
		}

		println(z)
	}
	return z, nil
}