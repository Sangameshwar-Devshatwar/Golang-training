package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Println("hellloo india") //prints the line as it is
	// str := "hello"
	// fmt.Printf("%d %s hello welcome to jumanji", 34, str) //this prints according to type of format and text also
	bytewritten, err := fmt.Fprint(os.Stdout, "golang\n") //reads the number of bytes and return the count
	if err != nil {
		panic(err)
	}
	fmt.Print(bytewritten, "byteswriiten\n")

	bytewritten1, err := fmt.Fprintf(os.Stdout, "%dthis is the Fprintf and example taken as a number is ", 234) // this function prints according
	if err != nil {
		panic(err)

	}
	fmt.Print(bytewritten1)

	var i int
	var str1 string

	r := strings.NewReader("hey how are you all")
	n, err := fmt.Fscan(r, &i, &str1)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fscanf:%v\n", err)
	}
	fmt.Printf("i:%v,i:%v,n:%v", n, i, str1)

}
