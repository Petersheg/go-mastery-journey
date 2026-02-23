package main

import (
	"fmt"
)

func main(){
	// Sum the integer from down to 1
	sum := SumUp(10)
	fmt.Println(sum)


	// Sume without the initial 
	sum_01 := SumUpWithoutInit(3)
	fmt.Println(sum_01)

	//For as while loop
	sum_02 := ForAsWhile()
	fmt.Println(sum_02)
}