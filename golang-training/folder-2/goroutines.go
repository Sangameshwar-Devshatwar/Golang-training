package main

import (
	"fmt"
	"time"
)

func main() {
	add(2, 4)
	time.Sleep(3 * time.Second)
	go sub(66, 54)
	time.Sleep(3 * time.Second)
	go mult(20, 30)
	div(30, 2)
}

func add(a int, b int) int {
	c := a + b
	fmt.Println("addition is", c)
	return c
}
func sub(a int, b int) int {
	c := a - b
	fmt.Println("substraction is", c)
	return c
}

func mult(a int, b int) int {
	c := a * b
	fmt.Println("multiplication  is", c)
	return c
}
func div(a int, b int) int {
	c := a / b
	fmt.Println("divison   is", c)
	return c
}
