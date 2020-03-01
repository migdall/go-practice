package main

import (
	"fmt"
)

func main() {
	x := 43 % 40
	y := 43 / 40
	fmt.Println(x, y)

	x = 0
	for {
		if x > 100 {
			break
		}

		if x % 2 == 0 {
			x++
			continue
		}

		fmt.Println(x)
		x++
	}
}
