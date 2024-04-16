package main

import (
	"fmt"
	"os"
)

func CreateFile() { // to reate a file
	f, err := os.Create("cityInfo.txt")
	if err != nil {
		panic(err)
	}
	data := []byte("mumbai is the capital of maharashtra")
	f.Write(data)
	f.WriteString("\ngandhinagar is the capital of Gujrat")
	defer f.Close()

}

func OpenFile() { //to open a file and read the data from file--type_1
	f, err := os.Open("cityInfo.txt")
	if err != nil {
		panic(err)
	}
	data := make([]byte, 100)
	f.Read(data)
	//fmt.Println(data)//this will print the file data in binary format soo we have to convert it into string
	fmt.Println(string(data))
}
func OpenFile2() { //by type_2 we don't need to open file just use (os.readfile)
	data, err := os.ReadFile("cityInfo.txt") //in data the text present in file will get stored soo need to read here data first
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
}
func main() {
	CreateFile()
	OpenFile()
	OpenFile2()
}
