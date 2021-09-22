package main

import "fmt"

func main() {
	var x [5]int // it's empty
	var y [6]int
	y[3] = 44 // assigning some value to the 4th block
	fmt.Println(x)
	fmt.Println(y)
}
