package main

import "fmt"

func main() {
	var word string
	fmt.Println("Enter a string for reversing")
	fmt.Scan(&word)
	reversed := ReverseString(word)
	fmt.Println("Reversed string is", reversed)
}
func ReverseString(str string) string {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed
}
