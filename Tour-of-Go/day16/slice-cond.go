package main

import (
	"fmt"
	"strings"
)

/* Creating a slice with make
Slices can be created with the built-in make function; this is how you create dynamically-sized arrays.

The make function allocates a zeroed array and returns a slice that refers to that array:

a := make([]int, 5)  // len(a)=5
To specify a capacity, pass a third argument to make:

b := make([]int, 0, 5) // len(b)=0, cap(b)=5

b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
*/ 
func MakeWithSlice() {
	a := make([]int, 5)
	fmt.Println("a", a)

	b := make([]int, 0, 5)
	fmt.Println("b", b)

	// Slice 2 elements from the beginning
	c := b[:2]
	printSlice("c", c)

	// Slice 3 elements from the second index
	d := c[2:5]
	printSlice("d", d)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len: %d, cap: %d, %v\n", s, len(x), cap(x), x)
}


/* Slices of slices
	Slices can contain any type, including other slices
*/ 
func TicTackToe() {
	// Create a Tic Tack Toe game
	board := [][]string{
		[]string {"_", "_", "_"},
		[]string {"_", "_", "_"},
		[]string {"_", "_", "_"},
	}

	board[0][0] = "x"
	board[2][2] = "0"
	board[1][2] = "x"
	board[1][0] = "0"
	board[0][2] = "x"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

/*Appending to a slice
	It is common to append new elements to a slice, and so Go provides a built-in append function. 
	The documentation of the built-in package describes append.

	func append(s []T, vs ...T) []T
	The first parameter s of append is a slice of type T, and the rest are T values to append to the slice.

	The resulting value of append is a slice containing all the elements of the original slice plus the provided values.

	If the backing array of s is too small to fit all the given values a bigger array will be allocated. 
	The returned slice will point to the newly allocated array.
*/ 
func AppendToSlice() {
	var s []int
	printSlice("", s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice("", s)

	//adding more than one elements
	s = append(s, 3,5,22)
	printSlice("",s)
}