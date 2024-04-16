package main

import "fmt"

func main() {
	details := User{"sangameshwar", "sangam@123", 9421379450, 25}
	fmt.Println("User details are :-", details)
	fmt.Printf("User details are :%+v\n", details) //%+v will print data as per naming convention given in user struct like key value pairs but we have to use printf to print
}

type User struct { //struct data of user
	Username string
	Email    string
	Mobile   int
	Age      int
}
