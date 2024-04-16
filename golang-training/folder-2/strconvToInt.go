package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) //this creates buffer to enter a string
	scanner.Scan()

	str, err := strconv.ParseInt(scanner.Text(), 32, 64)
	fmt.Println(err)
	fmt.Println(str)
	

}
