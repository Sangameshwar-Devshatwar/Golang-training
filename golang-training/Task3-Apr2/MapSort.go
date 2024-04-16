package main

import (
	"fmt"
	"sort"
)

func main() {
	sortMap := make(map[int]string)
	sortMap[45] = "orange"
	sortMap[84] = "banana"
	sortMap[65] = "apple"
	sortMap[55] = "guava"

	var arr = []int{}
	for key, _ := range sortMap {
		arr = append(arr, key)
	}
	sort.Ints(arr)
	for _, value := range arr {
		fmt.Printf("%v : %v\n", value, sortMap[value])
	}
}
