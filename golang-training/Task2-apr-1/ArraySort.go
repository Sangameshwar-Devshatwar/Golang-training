package main

import "fmt"

func main() {
	arr := []int{23, 45, 67, 89, 32, 45}
	fmt.Println("Elements before sorting are\n", arr)

	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}

		}

	}
	fmt.Println("Elements after sorting are::")
	fmt.Println(arr)

}
