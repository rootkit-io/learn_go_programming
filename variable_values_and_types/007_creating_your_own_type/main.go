package main

import "fmt"

var y int

type badboy int // ASSIGN badboy as TYPE that stores the value of INT

var z badboy

func main() {
	y = 32
	z = 49
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	// y = z // Cannot do like this because go is STATIC programming lang but there are ways we will discover soon.
	// fmt.Println(z)
	// fmt.Printf("%T\n", z)

	// But if you want to do it like that you can do something called CONVERSION in other programming lang it's \
	// called TYPECASTING but there is nothing like CASTING in go let's do it

	y = int(z)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

}
