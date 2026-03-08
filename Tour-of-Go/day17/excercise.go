package main

import (
	"fmt"
	"math/rand"
	// "golang.org/x/tour/pic"
)

/* Exercise: Slices

Implement Pic. It should return a slice of length dy, each element of which is a slice of dx 8-bit
unsigned integers. When you run the program, it will display your picture, interpreting the
integers as grayscale (well, bluescale) values.

The choice of image is up to you. Interesting functions include (x+y)/2, x*y, and x^y.

(You need to use a loop to allocate each []uint8 inside the [][]uint8.)

(Use uint8(intValue) to convert between types.)

*/ 

func Pic(dx, dy int) [][]uint8 {
	return CreateSlice(dx)
}

func CreateSlice(lent int) [][]uint8 {

	k := make([][]uint8, lent)

	for i := range k{
		k[i] = make([]uint8, lent)
	
		for j := range k[i] {
			k[i][j] = getRand(10)
		}
	}

	fmt.Println(k)
	return k
} 

func getRand(rang int) uint8 {
	x, y := rand.Intn(rang) + 1 , rand.Intn(rang) + 1

	random := []uint8{ uint8((x+y)/2), uint8(x*y), uint8(x^y) }
	
	
	return random[rand.Intn(len(random))]
}