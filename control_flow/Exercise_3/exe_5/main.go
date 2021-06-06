package main

import "fmt"

func main() {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}
		if x%4 != 0 {
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("done")
}
