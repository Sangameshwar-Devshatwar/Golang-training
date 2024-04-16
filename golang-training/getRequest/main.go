package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	fmt.Println("welcome to get request using golang")
	PerformGetRequest()
}

func PerformGetRequest() {
	const myurl = "https://dummy.restapiexample.com/api/v1/employees"
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("status code is", response.StatusCode)
	fmt.Println("status is", response.Status)
	fmt.Println("content length is", response.ContentLength)
	var contentstring strings.Builder
	result, _ := ioutil.ReadAll(response.Body)
	finalcontent, _ := contentstring.Write(result)
	fmt.Println("byte count is", finalcontent)
	fmt.Println(contentstring.String()) //string builder is used here to get string in response from that above url
	fmt.Println(string(result))
	//string builder is more useful than string because in small data we can use string but for big data we have to use string builder
}
