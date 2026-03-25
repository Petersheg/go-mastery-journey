package main

import (
	"fmt"
)

type I interface {
	M()
}

type T struct {
	S  string
}

/* Interface values with nil underlying values

	If the concrete value inside the interface itself is nil, the method will be called with a nil receiver.

	In some languages this would trigger a null pointer exception, but in Go it is common to write methods 
	that gracefully handle being called with a nil receiver (as with the method M in this example.)

	Note that an interface value that holds a nil concrete value is itself non-nil.
*/ 
func (t *T) M(){
	if t == nil {
		fmt.Println("<nil>")
		return
	}

	fmt.Println(t.S)
}

func NilInterface() {
	var i I
	var t *T

	i = t // nill value
	describe(i)
	i.M()

	i = &T{"Hello, world!"} // actual value
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v %t)\n", i, i)
}

/* Nil interface values
	A nil interface value holds neither value nor concrete type.

	Calling a method on a nil interface is a run-time error because there is no type 
	inside the interface tuple to indicate which concrete method to call.
*/
func NilValueInterface(){
	var i I
	describe(i)
	i.M() // panic: runtime error: invalid memory address or nil pointer dereference
}

/* The empty interface
	The interface type that specifies zero methods is known as the empty interface:

	interface{}
	An empty interface may hold values of any type. (Every type implements at least zero methods.)

	Empty interfaces are used by code that handles values of unknown type. 
	For example, fmt.Print takes any number of arguments of type interface{}.
*/ 
func EmptyInterface(){
	var i interface{}
	describeI(i)

	i = 42
	describeI(i)

	i = "Hello"
	describeI(i)
}

func describeI(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}