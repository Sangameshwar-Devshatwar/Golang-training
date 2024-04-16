package main

import (
	"test/folder1"
	"test/folder2"
)

func main() {
	// Call functions from folder1 package    folder1.FunctionFromFolder1()

	// Call functions from folder2 package
	folder1.FunctionFromFolder1()
	folder2.AnotherFunctionFromFolder2()
}
