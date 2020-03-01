package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's get to printing")
	// Begin with number 97 and end with number 122
	// Print out the number and the number's formatted ascii conversion
	for x := 32; x < 123; x++ {
		fmt.Printf("%d:%q\n", x, x)
	}
	fmt.Println("done")
}
