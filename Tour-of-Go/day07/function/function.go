package function

// A function can take zero or more arguments.
// In this example, add takes two parameters of type int.
// Notice that the type comes after the variable name.
func Add(x int, y int) int {
	return x + y
}

// When two or more consecutive named function parameters share a type, 
// all the type can be omitted but the last.
func AddShortened(x,y int) int {
	return x + y
};

// A function can return any number of results.
// The swap function returns two strings.
func Swap(x, y string) (string, string) {
	return y, x
}