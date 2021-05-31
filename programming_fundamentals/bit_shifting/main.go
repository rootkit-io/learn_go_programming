package main

import "fmt"

//func main() {
//x := 4
//fmt.Printf("%d\t\t%b\n",x,x)
//y := x << 1
//fmt.Printf("%d\t\t%b\n",y,y)
//}

// Now let's a look at kb , mb and gb in binary

//func main() {
//	kb := 1024
//	mb := 1024 * kb
//	gb := 1024 * mb
//
//	fmt.Printf("%d\t\t%b\n", kb, kb)
//	fmt.Printf("%d\t\t%b\n", mb, mb)
//	fmt.Printf("%d\t%b", gb, gb)
//
//}

// doing the above operation with the help of iota

const (
	_  = iota
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)
)

func main() {
	fmt.Printf("%d\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t%b\n", mb, mb)
	fmt.Printf("%d\t%b", gb, gb)
}
