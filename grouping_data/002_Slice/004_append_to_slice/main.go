package main

import "fmt"

func main() {
	x := []int{1,2,3,4,5}
	fmt.Println(x)
	x = append(x,6,7,8)
	fmt.Println(x)

	y := []int{69,79,89,99}
	x = append(x, y...)	// so what we have used a func. called varadic func.
	fmt.Println(x)
}
