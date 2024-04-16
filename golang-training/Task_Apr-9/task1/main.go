//	Task 1: Pointer Manipulation - Write a program that takes an integer input from the user and prints
//
// its memory address along with the value stored at that address.*
package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Println("Enter the input that you want ")
	fmt.Scan(&input)
	var num = &input
	fmt.Printf("the value at %v at adress is %d", num, *num)

}
