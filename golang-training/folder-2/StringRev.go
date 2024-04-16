package main

import "fmt"

func revString(str string) string {
	reversed := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversed += string(str[i])
	}
	return reversed
}
func main() {

	var str string = "welcome"
	fmt.Println(revString(str))


}
