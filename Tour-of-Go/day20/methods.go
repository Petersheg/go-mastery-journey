package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

/* Methods

	Go does not have classes. However, you can define methods on types.
	A method is a function with a special receiver argument.

	The receiver appears in its own argument list between the func keyword and the method name.
	In this example, the Abs method has a receiver of type Vertex named v.
*/ 
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}


func ReceiverMethod() {
	v := Vertex{X:3, Y:3}

	fmt.Println("Receiver Method", Abs(v))
}

/* Methods are functions
	Remember: a method is just a function with a receiver argument.

	Here's Abs written as a regular function with no change in functionality.
*/ 
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X * v.X + v.Y * v.Y)
}

func OrdinaryMethod() {
	v := Vertex{ X:5, Y:5 }

	fmt.Println("Ordinary Method", v.Abs())
	// fmt.Println(Abs(v))
}

/*
You can declare a method on non-struct types, too.

In this example we see a numeric type MyFloat with an Abs method.

You can only declare a method with a receiver whose type is defined in the same package as the method. 
You cannot declare a method with a receiver whose type is defined in another 
package (which includes the built-in types such as int).
*/ 
type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return  float64(-f)
	}

	return  float64(f)
}

func OtherTypeMethod() {
	f := MyFloat(-math.Sqrt2)

	fmt.Println("Method on other type", f.Abs())
}

