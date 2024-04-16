package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.youtube.com/watch?v=ru53LpdVHn4&list=PLRAV69dS1uWQGDQoBYMZWKjzuhCaOnBpa&index=25"

func main() {
	fmt.Println("handling web requests")
	response, err := http.Get(url) //it is doing get request to url by using http.Get from package "net/http".
	if err != nil {
		panic(err)

	}
	fmt.Printf("the type of response is %T\n", response)
	
	defer response.Body.Close() //this is user's responsibility to close the connection or response body

	bytes,err:=ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)

	}
	content:=string(bytes)
	fmt.Println(content)
}
