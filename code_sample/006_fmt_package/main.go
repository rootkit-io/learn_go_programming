package main

import "fmt"

var z = 420

func main() {
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Printf("%b\n", z)
	fmt.Printf("%x\n", z)
	fmt.Printf("%#x\n", z)
	fmt.Printf("%v\n", z)

	z = 911
	fmt.Printf("%#x\t%b\t%x\t%T\n", z, z, z, z)
	s := fmt.Sprintf("%#x\t%b\t%x\t%T\n", z, z, z, z)
	fmt.Println(s)
}

// This i've read the go documentation and implemented some code so
// i don't think so something is there to explain
// the link to doc is :(https://pkg.go.dev/fmt)

//keep going
