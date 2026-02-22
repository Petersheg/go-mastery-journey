package function

import "fmt"

func main() {
	fmt.Println(Add(2,5))

	//Function with a shortened parameter type
	fmt.Println(AddShortened(4,7))

	//function with double return types
	a, b := Swap("This", "That");
	println(a,b)
}