// array reversal
// Finding Maximum Element in Array - using in request values in int data type
// Sum of Array Elements
// Array Average Calculation
// Array Sorting
// Array Search

package main

import "fmt"

func main() {
	arr1 := []int{11, 45, 56, 87, 67, 90, 45}
	var revarr = make([]int, len(arr1))

	revarray(revarr, arr1)
}
func revarray(revarr []int, arr []int) {
	j := 0
	for i := len(arr) - 1; i >= 0; i-- {
		revarr[j] = arr[i]
		j++
	}
	fmt.Println(revarr)

}
