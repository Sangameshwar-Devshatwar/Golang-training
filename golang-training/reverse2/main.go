package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	empty := ""
	//sentence1 := ""
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter sentence you want to reverse every word")
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	sentence := strings.Split(input, " ")
	for _, val := range sentence {

		str1 := ReverseString(val)
		empty += " "
		empty += str1

	}

	fmt.Println("reversed string with every word reverse is", empty)

}
func ReverseString(str string) string {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		
		reversed += string(str[i])
	}
	return reversed
}
