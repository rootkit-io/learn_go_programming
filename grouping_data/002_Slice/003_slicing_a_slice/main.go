package main

import "fmt"

func main() {
	x := []int{4, 3, 2, 1, 7}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	// let's try to access all the slices with for loop without range
	// but we will first see with range.

	for i, v := range x {
		fmt.Println(i, v)
	}
	// now down below is the real deal.
	for i := 0; i <= 4; i++ {
		fmt.Println(i, x[i])
	}

}
