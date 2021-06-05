package main

import "fmt"

func main() {
	x := 42
	if x == 40 {
		fmt.Println("Our value was 40")
	} else if x == 41 {
		fmt.Println("Our value was 41")
	} else if x == 33 {
		fmt.Println("Our value was not 33")
	} else {
		fmt.Println("Our value was not 40, 41 and 33")
	}
}
