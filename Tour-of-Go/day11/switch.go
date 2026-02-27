package main

import (
	"fmt"
	"runtime"
	"time"
)

// Declaring a struct for Go runtime values
type RunTimeInfo struct {
	GOOS string
	GOARCH string
	Version string
	NumCPU int
	NumGoRoutine int
}

// Initiatiating RuntimeInfo with actual data.
var info RunTimeInfo = RunTimeInfo { 
	GOOS: runtime.GOOS,
	GOARCH: runtime.GOARCH,
	Version: runtime.Version(),
	NumCPU: runtime.NumCPU(),
	NumGoRoutine: runtime.NumGoroutine(),
};

/*
	A switch statement is a shorter way to write a sequence of if - else statements. 
	It runs the first case whose value is equal to the condition expression.
	Go's switch is like the one in C, C++, Java, JavaScript, and PHP, except that Go only runs the selected case, not all the cases that follow. 
	In effect, the break statement that is needed at the end of each case in those languages is provided automatically in Go. 
	Another important difference is that Go's switch cases need not be constants, and the values involved need not be integers.

*/
func SwitchOS() {
	fmt.Printf("Go runtime Objects %v \n", info)
	fmt.Print("Go runs on ")

	switch os := info.GOOS; os {
	case "darwin" :
		fmt.Println("MacOs. ")
	case "linux":
		fmt.Println("Linux. ")
	default :
		fmt.Printf("%s. \n", os)
	}
}

func SwitchDay() {
	fmt.Print("When is Saturday? ")

	today := time.Now().Weekday()
	
	switch today {
	case today + 0:
		fmt.Println("Today. ")
	case today + 1:
		fmt.Println("Tomorrow. ")
	case today + 2:
		fmt.Println("In two days. ")
	default:
		fmt.Println("Too far away. ")
	}
}

/*
	Switch without a condition is the same as switch true.
	This construct can be a clean way to write long if-then-else chains.
*/
func SwitchTrue() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning! ")
	case t.Hour() < 17:
		fmt.Println("Good afternoon")
	default:
		fmt.Println("Good evening!")
	}
}