package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "welcome to the world of go and if u want to know moew just visit go.in"
	file, err := os.Create("./myfile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	fmt.Println("lenght is", length)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	readfile("./myfile.txt")
}
func readfile(filename string) {
	reading, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println(reading)         //this will read the file but give the op of data in number coded format
	fmt.Println(string(reading)) //this will return total data that is into file into string format
}
