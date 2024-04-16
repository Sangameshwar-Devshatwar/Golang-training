package folder1

import (
	"fmt"
	"test/folder2"
)

func FunctionFromFolder1() {
	folder2.FunctionFromFolder2()
	fmt.Println("Function from folder1 called")
}
