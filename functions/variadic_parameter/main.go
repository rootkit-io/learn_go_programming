package main

import "fmt"

func main() {
	x := sum(2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("The total is,", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for the item at index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	//fmt.Println("The total sum is,", sum)
	return sum

}
