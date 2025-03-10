// Problem:
// TODO
// https://go.dev/tour/basics/13
// The expression T(v) converts the value v to the type T.

package main

import "fmt"

func main() {
	three := 3
	oneThird := float64(1.0 / 3.0)
	twoPlusThree := 2 + three
	// These two variables are of distinct types. To perform operations
	// between variables they must be of the same type! Convert one of
	// the variables in the Println so that the result of 1/3+2+3 is printed out.
	// Don't add lines or change the lines above!
	fmt.Println(oneThird + float64(twoPlusThree))
}
