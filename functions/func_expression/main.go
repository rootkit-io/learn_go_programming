package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("My first func expression")
	}
	f()

	g := func(x int) {
		fmt.Println("The year I born", x)
	}
	g(2005)
}
