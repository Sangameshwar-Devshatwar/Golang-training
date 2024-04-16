// package main

// import (
// 	"fmt"
// )

// func main() {
// 	// var no int
// 	// fmt.Scanln(&no)
// 	var number int
// 	fmt.Println("Enter number that you want for sum of specified elements of array arr= {89, 56, 5, 2, 9}")
// 	fmt.Scan(&number)
// 	var sum int = 0
// 	arr := []int{2, 7, 5, 11}
// 	for i := 0; i < len(arr); i++ {
// 		sum = sum + arr[i]
// 		//fmt.Println(sum)
// 	}
// 	for j := 0; j < len(arr); j++ {
// 		sum = sum - arr[j]
// 		if sum == number {
// 			fmt.Println(j, j+1)
// 			break

//			}
//		}
//	}
package main

import "fmt"

func main() {
	var nums []int = []int{3, 2, 4, 3, 5, 6, 7, 8, 8}
	target := 16

	twoSum(nums, target)
}

// func twoSum(nums []int, target int) []int {
// 	numMap := make(map[int]int)

// 	for i, num := range nums {
// 		complement := target - num
// 		fmt.Println(complement)
// 		if val, found := numMap[complement]; found {
// 			fmt.Println(found)
// 			fmt.Println(val)
// 			return []int{numMap[complement], i}
// 		}
// 		numMap[num] = i
// 	}

// 	return nil
// }

func twoSum(nums []int, target int) {
	numsmap := make(map[int]int)
	for i := 0; i < len(nums)-1; i++ {
		// if i+1 > len(nums) {
		// 	break
		// }
		result := nums[i] + nums[i+1]
		numsmap[i] = result
	}
	for num, _ := range numsmap {
		if numsmap[num] == target {
			fmt.Println(num, num+1)
		}
	}

}
