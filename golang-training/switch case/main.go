package main

import (
	"fmt"
	"math/rand"
)

func main() {
	diceNumber := rand.Intn(6) + 1 //random number gets generated here
	fmt.Println(diceNumber)
	switch diceNumber { //took random number as a switch case to use
	case 1:
		fmt.Println("dice value is 1")
	case 2:
		fmt.Println("dice value is 2")
		fallthrough
	case 3:
		fmt.Println("dice value is 3")
	case 4:
		fmt.Println("dice value is 4")
		fallthrough
	case 5:
		fmt.Println("dice value is 5")
	case 6:
		fmt.Println("dice value is 6")

	}
}
