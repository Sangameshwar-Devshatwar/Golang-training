package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Println("ENTER A NUMBER THAT YOU WANT:")
	fmt.Scanln(&input)
	if input%2 == 0 {
		fmt.Println("NO IS even ")
	} else if input == 0 {
		fmt.Println("0 cant be ven or odd")

	} else {
		fmt.Println("number is odd")
	}

}
