package main

import "fmt"

func main() {
	fmt.Println("inside main function")
	greet()
}
func greet() {
	fmt.Println("Inside greet function")
	c := 2 + 3
	fmt.Println(c)
}
