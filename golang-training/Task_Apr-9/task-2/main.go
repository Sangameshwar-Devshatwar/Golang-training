// Slice Operations - Write a program that
// takes a slice of integers as input and prints the sum of all the elements in the slice.
package main

import "fmt"

func main() {
	total := 0
	var size int
	fmt.Println("Enter size of slice you want to make")
	fmt.Scan(&size)
	var sum = make([]int, size)
	var input int
	fmt.Println("Enter numbers into slices")

	for i := 0; i < size; i++ {
		fmt.Scan(&input)
		sum[i] = input
	}
	for _, val := range sum {
		total = total + val
	}
	fmt.Printf("the sum of the given slice is %d \n", total)

}
