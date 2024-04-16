package main

import (
	"fmt"
)

func main() {
	dataChan := make(chan int, 3)
	var no int
	for i := 0; i < 3; i++ {
		fmt.Println("Enter number to entry it to channel")
		fmt.Scan(&no)
		dataChan <- no

	}
	close(dataChan)
	TotalSum(dataChan)

}

func TotalSum(data chan int) {
	// sum := 0
	for i := range data {
		//num := <-data
		// sum = sum + i
		fmt.Println(i)
	}
}
