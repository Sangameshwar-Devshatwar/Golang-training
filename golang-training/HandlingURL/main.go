package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=ghbj456ghb"

func main() {
	fmt.Println("welcome to url Handling in go")
	fmt.Println(myurl)

	result, _ := url.Parse(myurl)//it parses a raw url into url specific structure for further uses
	fmt.Println(result.Scheme)   //o/p:--https
	fmt.Println(result.Host)     //o/p:--lco.dev:3000
	fmt.Println(result.Path)     //o/p:--/learn
	fmt.Println(result.RawQuery) //o/p:--coursename=reactjs&paymentid=ghbj456ghb
	fmt.Println(result.Port())   //o/p:--3000

	qparams := result.Query()
	fmt.Printf("the type of qparams is\n", qparams)
	fmt.Println("\nthe value is:", qparams["coursename"])

	for val, value := range qparams {
		fmt.Printf("\nthe %v:%v", val, value)
	}
	partsofurl := &url.URL{ //it is of type struct format to define partsof url manually and makes,--
							//it into one independent website that when we get output we can open it by clicking on that generated site
		Scheme:  "https",
		Host:    "lco.dev",
		Path:    "/tutcss",
		RawPath: "user=hitesh",
	}
	fmt.Println(partsofurl) //o/p:--url is: https://lco.dev/tutcss
}
