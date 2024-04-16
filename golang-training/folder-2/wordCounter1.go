package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	counterWords := make(map[string]int)
	fmt.Println("Enter sentence that you want for word counter")
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	text := scanner.Text()
	words := strings.Fields(text)

	for _, word := range words {
		counterWords[word]++
	}
	for word, count := range counterWords {
		fmt.Println(word, count)
	}

}
