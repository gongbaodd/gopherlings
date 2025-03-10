//
// Problem:
// Go is a statically, strongly typed language:
//  - All values in go have a defined type at compile time (static typing)
//  - Operations between values of different types is not allowed (strong typing)
//
// Operations can only be done with values of the same type.
//
// This also means a type coercion does not happen in Go. For example: You can't
// add a boolean to a number without converting it with a function first.
//

package main

import (
	"fmt"
	"strconv"
)

func main() {
	// Oops! Looks like someone wanted to add 2.7 to Pi but
	// typed it in as a string. Change 2.7 to be a number so
	// that the result of 3.14159 plus 2.7 is printed.
	num, _ := strconv.ParseFloat("2.7", 64)
	fmt.Println(3.14159 + num)
}
