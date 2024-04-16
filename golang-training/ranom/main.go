package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	//rand.Seed(time.now().UnixNano())
	randno := rand.Intn(101)
	fmt.Println("random no between 1 to 101 is", randno)
	//**math usages**
	fmt.Println("sqrt of 25 is", math.Sqrt(25))
	//**max usages**
	fmt.Println("max between 10 and 20 is", math.Max(10, 20))

}
