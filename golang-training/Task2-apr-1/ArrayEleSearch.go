package main

import "fmt"

func main() {
	arr := []int{11, 45, 56, 87, 67, 90, 45}

	fmt.Println("enter number to search from given array :-{11, 45, 56, 87, 67, 90, 45} ")
	var Element int
	fmt.Scan(&Element)
	SearchEle(arr, Element)
}

func SearchEle(arr1 []int, val int) {
	for i := 0; i < len(arr1); i++ {
		if arr1[i] == val {
			fmt.Printf("element %d searched is found at index %d", arr1[i], i)
			break
		}
	}
	fmt.Println("element does not exists into array given by user")

}
