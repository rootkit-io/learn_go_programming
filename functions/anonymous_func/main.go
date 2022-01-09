package main

import "fmt"

func main() {
	foo()
	func() {
		fmt.Println("Hii from anon func")
	}()

	func(x int) {
		fmt.Println("The Meaning of life", x)
	}(42)
}

func foo() {
	fmt.Println("foo ran sucessfully yay!!!!!")
}
