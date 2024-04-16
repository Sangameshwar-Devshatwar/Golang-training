package main

func main() {
	// 	var target int
	// 	fmt.Println("enter target of two sum that you want to set")
	// 	fmt.Scan(&target)
	// 	var sum int
	// 	nums := []int{2, 7, 11, 15}
	// 	for i := 0; i < len(nums); i++ {
	// 		for j := 0; j <= i; j++ {
	// 			sum = sum + nums[j]
	// 			if sum == target {
	// 				// arr := []int{i, j}
	// 				fmt.Println(i+1, j+1)
	// 				break

	// 			}
	// 			sum = 0

	// 		}

	// 	}
	// }
	for i := 0; i < len(nums); i++ {
		// var elements [3]int
		j := 1

		sum := nums[i] + nums[j]
		if sum == target {
			return []int{i, j}
		}
		sum = 0
		j++

	}
	return nil
}
