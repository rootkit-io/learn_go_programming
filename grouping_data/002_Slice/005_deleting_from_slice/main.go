package main

import "fmt"

func main() {
	x := []int{11, 22, 33, 44, 55, 66, 4, 6, 7, 8}
	fmt.Println(x)

	// we dont need a new function for deleting a slice we can do with append.

	x = append(x[:4], x[6:]...)
	fmt.Println(x)
}
