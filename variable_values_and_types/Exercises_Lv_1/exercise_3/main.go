package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v\t", x, y, z) // %v for print the value and \t for space
	fmt.Println(s)
}
