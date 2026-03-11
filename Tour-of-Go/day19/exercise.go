package main

import "fmt"

/* Exercise: Fibonacci closure

Let's have some fun with functions.

Implement a fibonacci function that returns a function (a closure)
that returns successive fibonacci numbers (0, 1, 1, 2, 3, 5, ...).
*/
func fibonacci() func() int {
	a, b := 0, 1

	return func() int {
		result := a
		a, b = b, a+b

		return result
	}
}

func Fibonacci() {
	f := fibonacci()
	
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}