package main

import "fmt"

/*
An array has a fixed size. A slice, on the other hand, is a dynamically-sized,
flexible view into the elements of an array. In practice, slices are much more common than arrays.

The type []T is a slice with elements of type T.

A slice is formed by specifying two indices, a low and high bound, separated by a colon:

a[low : high]
This selects a half-open range which includes the first element, but excludes the last one.

The following expression creates a slice which includes elements 1 through 3 of a:

a[1:4]

*/ 
func Slice() {
	primes := []int{ 1, 3, 5, 7, 11, 13 }

	var s []int = primes[1:4]
	fmt.Println(s)
}


/* Slices are like references to arrays

	A slice does not store any data, it just describes a section of an underlying array.
	Changing the elements of a slice modifies the corresponding elements of its underlying array.

	Other slices that share the same underlying array will see those changes.

*/ 
func SliceIsAReference() {
	names := [4]string{"John", "Jame", "Okorocha", "Ezekwezeli"}
	fmt.Println(names)
	arrlen := len(names)

	// Create 2 slice with 2 elements in each
	a := names[0:2]
	b := names[1:arrlen] 
	fmt.Println(a,b)

	b[0] = "Ademeko"
	fmt.Println(a,b)
	fmt.Println(names)
}

/* Slice literals

	A slice literal is like an array literal without the length.

	This is an array literal:

	[3]bool{true, true, false}
	And this creates the same array as above, then builds a slice that references it:

	[]bool{true, true, false}
*/ 
func SliceLiteral() {
	q := []int{2,4,5,6,7,9}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct { X int; Y bool } {
		{ 2, true },
		{ 4, false },
		{ 5, true },
		{ 6, true },
		{ 7, false },
		{ 9, true },
	}

	fmt.Println(s)
}

/* Slice defaults

When slicing, you may omit the high or low bounds to use their defaults instead.

The default is zero for the low bound and the length of the slice for the high bound.

For the array

var a [10]int
these slice expressions are equivalent:

a[0:10]
a[:10]
a[0:]
a[:]
*/ 
func SliceDefault() {
	n := []int{ 3, 7, 9, 11, 13 }

	// Prints 2 elements from index 1 to 2
	k := n[1:3]
	fmt.Println(k)

	// Prints 3 element from first index to 2
	l := n[:3]
	fmt.Println(l)

	// Prints 4 elements from index 1 to slice len
	m := n[1:]
	fmt.Println(m)

}


/* Slice length and capacity

A slice has both a length and a capacity.

The length of a slice is the number of elements it contains.

The capacity of a slice is the number of elements in the underlying array, 
counting from the first element in the slice.

The length and capacity of a slice s can be obtained using the expressions len(s) and cap(s).

You can extend a slice's length by re-slicing it, provided it has sufficient capacity. 
Try changing one of the slice operations in the example program to extend it beyond its 
capacity and see what happens.

*/ 
func SliceLengthAndCapacity() {
	s := []int {3, 5, 7, 11, 13}
	printSlices(s)

	// Slice the slice to give it zero length
	s = s[:0]
	printSlices(s)

	// Extend it length
	s = s[:4]
	printSlices(s)

	// Drop the first two values
	s = s[2:]
	printSlices(s)
}

func printSlices(s []int) {
	fmt.Printf("len=%d, cap=%d %v\n", len(s), cap(s), s)
}

/* Nil slices

	The zero value of a slice is nil.
	A nil slice has a length and capacity of 0 and has no underlying array.
*/ 
func NilSlice() {
	var s []int
	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("Nil!")
	}

}
