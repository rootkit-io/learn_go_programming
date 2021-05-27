package main

import (
	"fmt"
)

func main() {
	s := "hello pews"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	bj := []byte(s)
	fmt.Println(bj)
	fmt.Printf("%T\n", bj)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}
	fmt.Println("")
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}

// best resource to learn more about STRINGS --> https://blog.golang.org/strings
