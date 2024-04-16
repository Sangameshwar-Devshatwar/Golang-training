package main

import "fmt"

func main() {
	var number int
	fmt.Println("Enter the number to check factorial")
	fmt.Scan(&number)
	fact := 1
	for i := 1; i <= number; i++ {
		fact = fact * i
	}
	fmt.Println("factorial of %v is %v", number, fact)
}
