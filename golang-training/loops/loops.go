package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		var sum = i + i + 1
		fmt.Println(sum)
	}
}
