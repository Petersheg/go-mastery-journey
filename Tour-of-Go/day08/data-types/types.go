package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	ToBe bool = false
	Maxint uint64 = 1<<64 - 1
	z complex128 = cmplx.Sqrt(-5 + 12i)
)

/*
	Go Basic Data Types
	1. bool
	2. string
	3. int, int8, int16, int32, int64
	4. uint, uint8, uint16, uint32, uint64, uintptr
	5. byte // alias for uint8
	6. rune // alias for int32
	7. float32, float64
	8. complex64, complex128
*/
/*
	The int, uint, and uintptr types are usually 32 bits wide on 32-bit systems and
	64 bits wide on 64-bit systems. When you need an integer value you should use int
	unless you have a specific reason to use a sized or unsigned integer type.
*/
func Types() {
	fmt.Printf("Type: %T, Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T, Value: %v\n", Maxint, Maxint)
	fmt.Printf("Type: %T, Value: %v\n", z, z)
}

/*
 	Variables declared without an explicit initial value are given their zero value.

	The zero value is:

	0 for numeric types,
	false for the boolean type, and
	"" (the empty string) for strings.
*/ 
func ZeroValues() {
	var i int
	var f float32
	var b bool
	var s string

	fmt.Printf("%v, %v, %v, %q\n", i, f, b, s) // %q is use for string formating.
}

//The expression T(v) converts the value v to the type T.
func TypeConversion() {
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y)) //converts int to float64

	var z uint = uint(f) //converts float64 to uint
	fmt.Println(x, y, z)
}

/*
	When declaring a variable without specifying an explicit type (either by using the := syntax or var = expression syntax),
	the variable's type is inferred from the value on the right hand side.
*/ 
func TypeInference() {

	//When the right handside is typed, the new variable is of the same type
	var i int
	j := i // J becomes int

	//But when the righ handside is untyped numeric, the new variable me be int, float64 or complex128 depending on the
	//precision of the constant.

	k := 3 //int
	l := 3.142 //float64
	g := 0.867 + 0.5i // complex128

	fmt.Printf( "%T, %T, %T, %T", j, k, l, g ) //Print the Type of the variable.
}

/*
	Constants are declared like variables, but with the const keyword.
	Constants can be character, string, boolean, or numeric values.
	Constants cannot be declared using the := syntax.
*/
const Pi = 3.142
func Constant() {
	fmt.Printf("This is a constant variable %v (%T) ", Pi, Pi)
}

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)
func needInt(x int) int {
	var k int = x*10 + 1
	return k 
}

func needFloat(x float64) float64 { 
	var j float64 = x * 0.1
	return  j
}

/*
	Numeric constants are high-precision values.
	An untyped constant takes the type needed by its context.
*/
func NumericConstant() {
	s_int := needInt(Small)
	s_float := needFloat((Small))
	b_float := needFloat(Big)

	fmt.Println(s_int, s_float, b_float)
}