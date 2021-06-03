package main

import "fmt"

func main() {
	x := 1
	for x < 10 { // For statement with single condition
		fmt.Println(x)
		x++
	}
	fmt.Println("Done.")
}
