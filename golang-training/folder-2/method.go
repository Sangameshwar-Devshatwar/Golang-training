package main

import "fmt"

func main() {
	details := User{"sangameshwar", "sangam@123", 9421379450, 25}
	fmt.Println("User details are :-", details)
	fmt.Printf("User details are :%+v\n", details) //%+v will print data as per naming convention given in user struct like key value pairs but we have to use printf to print
	details.getUsername()
	details.getEmail()
	details.getMobileno()
	details.getAge()
}

type User struct { //struct data of user
	Username string
	Email    string
	Mobile   int
	Age      int
}

////using struct.go example to see the use cases of methods in golang

func (u User) getUsername() { //method is created here,by creating obj of user struct as u
	fmt.Println("the username is", u.Username)
}
func (u User) getEmail() { //method is created here,by creating obj of user struct as u
	fmt.Println("the Email is", u.Email)
}
func (u User) getMobileno() { //method is created here,by creating obj of user struct as u
	fmt.Println("the mobileNo is", u.Mobile)
}
func (u User) getAge() { //method is created here,by creating obj of user struct as u
	fmt.Println("the Age is", u.Age)
}
