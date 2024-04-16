package main

import (
	"fmt"
)

func main() {
	var a = 500
	var b = 100
	if a > b {
		fmt.Println("a is greater than b")
	} else if b < a {
		fmt.Println("b is lesser than a")
	} else {
		fmt.Println("all neutral")
	}
}
