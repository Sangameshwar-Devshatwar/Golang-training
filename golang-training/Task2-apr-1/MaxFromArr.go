package main

import "fmt"

var Max int = 0

func main() {

	maxArr := []int{32, 45, 65, 67, 43, 67}
	maxarr(maxArr)

}

func maxarr(maxArr []int) {
	for i := 0; i < len(maxArr); i++ {
		if i >= len(maxArr)-1 {
			break
		}
		Max = max(maxArr[i], maxArr[i+1])
	}
	fmt.Println("maximum value from givenarray is:", Max)
}
