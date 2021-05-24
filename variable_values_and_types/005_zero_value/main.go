package main

import (
	"fmt"
)

var z string
var y int

func main() {
	// DECLARE the variable to be of certain TYPE
	// and then ASSIGN a VALUE of that variable
	fmt.Println("the value of z", z, "ending")
	fmt.Printf("%T\n", z)
	z = "007 james bond"
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	fmt.Println("the value of y", y)
	fmt.Printf("%T\n", y)
}

// false for booleans
// 0 	 for int
// 0.0   for float
// "" 	 for strings
// nil   for
//			* pointers
//			* functions
//			* interfaces
//			* slices
//			* channels
//			* maps

// TIPS: use short declaration as much as possible
// use var for zero value and package level scope
