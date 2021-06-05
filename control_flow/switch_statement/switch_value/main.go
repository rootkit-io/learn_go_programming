package main

import (
	"fmt"
)

func main() {
	//switch "Bond_James" { // this is one way
	//case "gold_digger":
	//	fmt.Println("miss money")
	//case "Bond_James":
	//	fmt.Println("yoyo this will print")
	//}

	// Another way

	x := "bond" // this is another way that you can define var
	switch x {  // and put it here
	case "boom", "007", "helew": // and you can put multiple here
		fmt.Println("boom boom")
	case "bond":
		fmt.Println("will this print? Yup this will ")
	}
}
