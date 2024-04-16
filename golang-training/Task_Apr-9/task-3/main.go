// Task 3: Array Operations - Write a program that initializes an array of integers
// and finds the maximum and minimum values in the array.
package main

import "fmt"

func main() {
	var arr [5]int
	var num int
	var size int
	fmt.Println("Enter size of slice you want to make")
	fmt.Scan(&size)
	fmt.Println("Enter elements to add in array")

	for i := range size {
		fmt.Scan(&num)
		arr[i] = num
	}

	for j := 0; j < len(arr); j++ {
		if j == len(arr)-1 {
			break
		}
		if arr[j] > arr[j+1] {
			temp := arr[j]
			arr[j] = arr[j+1]
			arr[j+1] = temp
		}

	}
	fmt.Printf("minimum value in array is %d:", arr[0])
	fmt.Printf("maximum value in array is %d:", arr[len(arr)-1])
}
