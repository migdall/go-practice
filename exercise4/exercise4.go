package main

import (
	"fmt"
)

type happy int

var x happy

func main() {
	fmt.Println(x)
	result := fmt.Sprintf("%T", x)
	fmt.Println(result)
	x = 42
	fmt.Println(x)
}
