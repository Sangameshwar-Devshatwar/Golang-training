package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println("fizz")
			if i%5 == 0 {
				fmt.Println("Buzz")
				if i%3 == 0 && i%5 == 0 {
					fmt.Println("fizzzzzbuzzzzzzz")
				}
			}
		}
	}
}
