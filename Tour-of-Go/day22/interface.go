package main

import (
	"fmt"
	"math"
)


type Abser interface{
	Abs() float64
};

/* Interfaces
	An interface type is defined as a set of method signatures.
	A value of interface type can hold any value that implements those methods.

	Note: There is an error in the example code on line 22. 
	Vertex (the value type) doesn't implement Abser because the Abs method is defined only on *Vertex (the pointer type).
*/
func Interface() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// This is an  intentional error to show that only *Vertex implements Abser, not Vertex. --- IGNORE ---
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

type Vertex struct{
	X, Y float64
}

func (v *Vertex) Abs() float64{
	return  math.Sqrt(v.X*v.X + v.Y*v.Y)
}



/* Interfaces are implemented implicitly
	A type implements an interface by implementing its methods. 
	There is no explicit declaration of intent, no "implements" keyword.

	Implicit interfaces decouple the definition of an interface from its implementation, 
	which could then appear in any package without prearrangement.
*/

type I interface {
	M()
}

type T struct {
	S string
}

/*This meathod means type T implement the interface I, but we dont need to explicitly declare that it does so */ 
func (t T) M() {
	fmt.Println(t.S)
}

func ExplicitInterface() {
	// initiate Type T
	var i I = T{ S : "Hello" } // OR i := T{"Hello"}

	i.M()
}




/* Interface values
	Under the hood, interface values can be thought of as a tuple of a value and a concrete type:

	(value, type)
	An interface value holds a value of a specific underlying concrete type.

	Calling a method on an interface value executes the method of the same name on its underlying type.
*/ 
type J interface {
	N()
}

type K struct {
	S string
}

func (k *K) N() {
	fmt.Println(k.S)
}

type F float64

func (f F) N(){
	fmt.Println(f)
}


func InterfaceValue() {
	var i J //Declaration

	i = &K{"Hellp"} //Initializing
	describe(i)
	i.N()

	i = F(math.Pi)
	describe(i)
	i.N()

}

func describe(i J) {
	fmt.Printf("(%v , %t)\n", i, i)
}
