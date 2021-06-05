package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print2")
	case (3 == 3):
		fmt.Println("prints")
		fallthrough // if we do not put this program will end and exits with a true value.
	case (4 == 4):
		fmt.Println("also true, does it print's ?")
		fallthrough
	case (7 == 9):
		fmt.Println("not true 1")
		fallthrough
	case (11 == 14):
		fmt.Println("not true 2")
		fallthrough
	case (15 == 15):
		fmt.Println("true 15")
		fallthrough
	default: // default case
		fmt.Println("this is default")

	}
}
