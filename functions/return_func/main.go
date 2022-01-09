package main

import "fmt"

func main() {
	fmt.Println(bar()())

}

func bar() func() int {
	return func() int {
		return 420
	}
}
