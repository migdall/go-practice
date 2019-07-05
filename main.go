package main

import (
	"fmt"
)

func main() {
	var x int = 23
	fmt.Printf("%T\n", x)

	// Print the result from the function
	// getName
	fmt.Println(getName())
}

func getName() string {
	return "Jesus"
}
