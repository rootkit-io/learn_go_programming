package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("001")
	}

	if false {
		fmt.Println("002")
	}

	if !true {
		fmt.Println("003")
	}

	if !false {
		fmt.Println("004")
	}

	if 2 == 2 {
		fmt.Println("005")
	}

	if 2 != 2 {
		fmt.Println("006")
	}

	if !(2 == 2){
		fmt.Println("007")
	}
	if !(2 != 2){
		fmt.Println("008")
	}
}
