package main

import "fmt"

var x bool

func main() {
	fmt.Println(x)
	x = true
	fmt.Println(x)
}


// read more about bool (https://golang.org/ref/spec#Boolean_types)
