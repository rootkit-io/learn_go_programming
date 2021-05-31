package main

import "fmt"

//const a = 45
//const b = 49.59
//const c = "Jesse Pinkman"

// but wait there is one more way from CONST if writing CONST is too boring then
// do it like this

//const (
//	a = 45
//	b = 49.5
//	c = "Jesse Pinkman"
//)

// and if you want to make a typed CONST then you can do like this in
// these are called TYPED CONST

const (
	a int     = 45
	b float64 = 49.5
	c string  = "Jesse Pinkman"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
