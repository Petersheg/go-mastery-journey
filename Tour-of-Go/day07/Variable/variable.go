package variable

import "fmt"

//The var statement declare a list of variables, as in function arguement lsit, the type is last.
//The var statement can be in the package level or in function level.

var java, c, python bool

func Vary() {
 var i int
 fmt.Println(i, c, python, java)
}

// A var declaration can include initializers, one per variable.
// If an initializer is present, the type can be omitted; 
// the variable will take the type of the initializer.

var i, j int = 2, 5
func VaryWithInitializer() {
	var c, python, java = true, false, "no" // Note that the declaration is without specific type.

	fmt.Println(i, j, c, python, java)
}


// Inside a function, the := short assignment statement 
// can be used in place of a var declaration with implicit type.
// Outside a function, every statement begins with a keyword 
// (var, func, and so on) and so the := construct is not available.

func VaryShorthand() {
	var i, j int = 3, 8
	k := 9
	c, python, java := true, false, "no"

	fmt.Println(i, j, k, c, python, java)

}