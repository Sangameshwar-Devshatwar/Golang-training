package main

import (
	"fmt"
)

func main() {
	var a int
	var b int
	var t int
	fmt.Println("Enter 2 nos for multiplication")
	fmt.Scanln(&a, &b)
	t = a * b
	fmt.Print("multiplication=", t)
}
