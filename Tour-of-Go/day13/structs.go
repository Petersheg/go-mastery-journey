package main

import "fmt"

/*
	Struct is a collection of fields

*/ 
type Vertex struct {
	X int
	Y int
}

//Print the declared Struct
func Structs() {
	fmt.Println(Vertex{3, 5})
}

// Sreuct fields are access using dot .
func AccessStruct() {
	v := Vertex{ 8, 4 }

	//Reassing X
	v.X = 20

	//Pring X new value
	fmt.Println(v.X)
}

/*	Pointers to Structs

	Structs fields can be access via struct pointer

	To access the field X of a struct when we have the struct pointer p we could write (*p).X. 
	However, that notation is cumbersome, so the language permits us instead to write just p.X, 
	without the explicit dereference.
*/
func StructPointers() {
	v := Vertex{ 90, 32 }
	p := &v

	p.X = 300
	p.Y = 20

	fmt.Println(v)
}

/*
	Struct Literals

	A struct literal denotes a newly allocated struct value by listing the values of its fields.

	You can list just a subset of fields by using the Name: syntax. (And the order of named fields is irrelevant.)
	The special prefix & returns a pointer to the struct value.

*/ 
func StructLiteral() {
	var (
		v1 = Vertex{2,5} //has type Vertex
		v2 = Vertex{X: 3} //Y is implicit
		v3 = Vertex{} // X:0, Y:0
		p = &Vertex{1,2} // has type *Vertex 
	)

	fmt.Println(v1, p, v2, v3)
}