package main

import "fmt"

// DECLARE the variable "y"
// ASSIGN the value 43
// declare & assign = initilization

var y = 433

// DECLARE with the VARIABLE with the IDENTIFIER "z"
// and that the VARIABLE with the IDENTIFIER "z" is of TYPE int
// ASSIGNS the ZERO VALUE of TYPE int to "z"
// FALSE for booleans, 0 for integers, 0.0 for floats, "" for strings,
// nil for pointers, functions, interfaces, slices, channels and maps.
var z int

func main() {
	// Short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 442
	fmt.Println(x)

	fmt.Println(y)

	foo()

	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}
