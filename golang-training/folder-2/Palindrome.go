package main

import (
	"fmt"
)

func main() {
	fmt.Println("Enter a  word to check is it palindrome or not")
	var word string
	fmt.Scan(&word)
	ans := checkPalindrome(word)
	//fmt.Println(ans)
	if ans == true {
		fmt.Println("The word is palindrome")

	} else {
		fmt.Println("the word is not palindrome")
	}
}

func checkPalindrome(word string) bool {
	j := len(word) - 1
	i := 0
	for i < j {

		if word[i] != word[j] {
			return false
		}
		i++
		j--

	}
	return true
}

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter input string:")
// 	input, _ := reader.ReadString('\n')
// 	result := strings.TrimSpace(input)

// 	if checkPalindrome(result) {
// 		fmt.Println("String is Palindrome.")
// 	} else {
// 		fmt.Println("String is not Palindrome.")
// 	}
// }

// func checkPalindrome(in string) bool {

// 	lowerCaseString := strings.ToLower(in)

// 	for i := 0; i < len(lowerCaseString)/2; i++ {

// 		// if lowerStr[i] != lowerStr[len(lowerStr)-1-i]
// 		if lowerCaseString[i] != lowerCaseString[len(lowerCaseString)-1-i] {
// 			return false
// 		}
// 	}
// 	return true
// }
