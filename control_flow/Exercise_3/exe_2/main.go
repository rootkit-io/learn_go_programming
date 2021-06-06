package main

import "fmt"

func main() {
	for i := 65; i <= 90; i++ {
		fmt.Println(i)
		for x := 0; x < 3; x++ {
			fmt.Printf("\t\t%#U\n", i)
		}
	}
}
