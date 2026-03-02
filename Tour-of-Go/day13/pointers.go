package main

import "fmt"

/*
	A Pointer holds memory address of a value.
	A type *T is a point to the value T, it zero value is nil.

	The & operator generates a pointer to its operand.
	i := 42
	p := &i

	The * operator denotes the pointer's underlying value.

	fmt.Println(*p) // reads i through the point P
	*p = 21 //Set i through the pointer P

	This is known as "dereferencing" or "indirecting".

	Summary
	& is use to SET a pointer
	* is use to READ Pointer's value.

*/ 
func Pointers() {
	i, j := 42, 2701

	p := &i //Points to i
	fmt.Println(*p) //Reads i throuhg the point P.

	*p = 21 //Set the value of i throuhg the pointer P
	fmt.Println(i) //Print =s the new value of i

	k := &j // Points to j
	*k = *k / 37 //Divides J through the point 

	fmt.Println(j) //Prints the new value of J.
}
