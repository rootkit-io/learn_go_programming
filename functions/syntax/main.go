package main

import (
	"fmt"
)

func main() {
	foo()
	bar("Jesse")
	s1 := woo("DonnerBoner")
	fmt.Println(s1)
	a, b := bang("Jesse ", "Pinkman")
	fmt.Println(a)
	fmt.Println(b)
}

func foo() {
	fmt.Println("hello from foo")
}

// Everything in Go is PASS by VALUE

func bar(s string) {
	fmt.Println("hello,", s)
}

func woo(s string) string {
	return fmt.Sprint("Hello from woo, ", s)
}

func bang(fn string, ln string) (string, bool) {
	z := fmt.Sprint(fn, ln, `, says "Hello",`)
	b := false
	return z, b
}

// func (r receiver) ) identifiers(parameters) (returns) {...}
