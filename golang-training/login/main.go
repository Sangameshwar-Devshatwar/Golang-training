package main

import (
	"fmt"
	"login/folder1"
)

var myUsername = "Sangam@123"
var myPass = "1234"
var username string
var password string

func main() {

	//var myPass = "1234"
	fmt.Println("Enter username")
	fmt.Scan(&username)
	fmt.Println("Enter password")
	fmt.Scan(&password)
	folder1.Login(username, password)

}

//	func reveal_password() {
//		fmt.Println("correct password is ", myPass)
//	}
