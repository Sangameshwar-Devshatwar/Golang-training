package main

import "fmt"

func main() {
	var countryList [4]string //declared array of length 4 of type String
	countryList[0] = "India"
	countryList[1] = "America"
	countryList[2] = "Russia"
	countryList[3] = "United Kingdom"
	fmt.Println("list of countries are as follows", countryList)
	fmt.Println("length of the coutrylist array is ", len(countryList))       // len() is a inbuild func to get the length of given array
	var cityNames = [5]string{"latur", "nanded", "pune", "bombay", "Lonavla"} //we can declare array like this also
	fmt.Println("the list of the cities are as", cityNames)
}
