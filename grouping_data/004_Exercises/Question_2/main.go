/* What is the length of a slice created using make([]int, 3, 9)? */
// The length should be 3 so let's make a program with the same specs and then see.

package main

import "fmt"

func main() {
	x := make([]int, 3, 9)
	fmt.Println(x)
}

