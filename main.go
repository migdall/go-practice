package main

import (
	"fmt"
)

type star string

var molly star

func main() {
	var x int = 23
	fmt.Printf("%T\n", x)

	s := fmt.Sprintf("%T", x)
	fmt.Println(s)

	// Print the result from the function
	// getName
	fmt.Println(getName())

	printStar()
}

func getName() string {
	return "Jesus"
}

func printStar() {
	fmt.Println(molly)
}
