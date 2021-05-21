package main

import "fmt"

var y = 42

// DECLARE the VARIABLE with the IDENTIFIER "z"
// is of TYPE string
// and I ASSIGN the value "Shaken, not stirred"

var z string = "shaken, not stirred"

var a string = `Jesse said,
"Shaken,

not stirred"`

// this is a STATIC programming language
// a VARIABLE is DECLARED to hold the VALUE of a certain TYPE
// not a DYNAMIC programming language

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	//z = 43
	//fmt.Println(z)
	//fmt.Printf("%T\n", z)
}



// keep going
