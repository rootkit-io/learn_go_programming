/* Given the following array, what would x[2:5] give you?
   x := [6]string{"a","b","c","d","e","f"}

Answer --> So this is going to give c,d,e.

Let's get this in the program and let's see.
 */

package main

import "fmt"

func main() {
	x := [6]string{"a","b","c","d","e","f"}
	fmt.Println(x[2:5])
}