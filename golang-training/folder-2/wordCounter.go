package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	wordCounts := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text (press Ctrl+D to end):")
	for scanner.Scan() {
		text := scanner.Text()
		words := strings.Fields(text)
		for _, word := range words {
			wordCounts[word]++
		}
	}

	fmt.Println("Word Counts:")
	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
}
