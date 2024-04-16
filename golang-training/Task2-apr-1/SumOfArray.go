package main

import "fmt"

// var sum int=0
func main() {
	var sum int = 0
	array := []int{23, 56, 78, 90, 45, 7, 8, 76}
	for i := 0; i < len(array); i++ {
		sum = sum + array[i]

	}
	fmt.Println("sum of array is", sum)
}
