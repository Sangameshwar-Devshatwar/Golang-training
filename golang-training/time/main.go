package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now()
	var name string
	var inputDate string
	fmt.Println("enter your name")
	fmt.Scan(&name)
	fmt.Println("Enter proper date format as: yyyy-dd-mm ")
	fmt.Scan(&inputDate)

	date, error := time.Parse("2006-01-02 ", inputDate)
	if error != nil {
		fmt.Println(error)
		return
	}

	///if currentTime.Hour()
	if currentTime.Hour() == 24 && currentTime.Hour() < 12 {
		fmt.Println("welcome to golang", name, "Good morning", currentTime.Format(" 15:04:05 Monday"), date)
	} else if currentTime.Hour() >= 12 && currentTime.Hour() < 17 {
		fmt.Println("welcome to golang", name, "Good Afternoon", currentTime.Format(" 15:04:05 Monday"), date)
	} else {
		fmt.Println("welcome to golang", name, "Good Night", currentTime.Format(" 15:04:05 Monday"), date)
	}

}
