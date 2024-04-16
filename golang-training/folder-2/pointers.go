package main

import "fmt"

func main() {
	muNumber := 23
	var ptr = &muNumber
	var numptr = &ptr
	fmt.Println(ptr)
	fmt.Println(*ptr)
	*ptr = *ptr * 2
	fmt.Println(*ptr)
	fmt.Println(*numptr)
	var p *int
	var num = 12
	p = &num
	fmt.Println(p)
	fmt.Println(*p)
	var n = &p
	fmt.Println(**n)
}
