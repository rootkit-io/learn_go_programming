package main

import "fmt"

var x int = 69
func main() {
	fmt.Printf("%d\t%b\t\t%#x\t\n",x,x,x)

	z := x << 1
	fmt.Printf("%d\t%b\t%#x\n",z,z,z)

}
