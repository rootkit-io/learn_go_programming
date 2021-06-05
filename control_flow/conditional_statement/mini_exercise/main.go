package main

import "fmt"

// This is just me revisiting control flow..
func main() {
	for x := 0; x < 100; x++ {
		if x%2 == 0 {
			fmt.Println(x)
		}
	}
}
