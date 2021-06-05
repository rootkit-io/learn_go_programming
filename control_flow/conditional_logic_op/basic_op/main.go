package main

import (
	"fmt"
)

func main() {
	fmt.Println(true && true)
	fmt.Println(true && false) // this is "END"
	fmt.Println(true || true)  // this is "OR"
	fmt.Println(true || false)
	fmt.Println(!true) // this will "REVERSE" the value like if it true it will make it false.
}
