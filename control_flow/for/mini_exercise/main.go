package main

import "fmt"

func main() {
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v corresponds to %+q in ASCII\n", i, i)
		i++
	}
}
