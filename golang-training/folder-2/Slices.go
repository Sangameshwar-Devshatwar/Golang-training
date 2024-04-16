package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"mango", "banana", "peach", "coconut", "pineapple"}
	fmt.Println("fruitlis is", fruitList)
	fruitList = append(fruitList[1:])
	fmt.Println("fruitlis is", fruitList)
	fruitList = append(fruitList, "guava", "apple", "watermelon") // this will append values at the end of fruitList
	fmt.Println("fruitlis is", fruitList)
	fruitList = (fruitList[3:])
	fmt.Println("fruitlis is", fruitList)
	fruitList = append(fruitList[2:4]) //this will return values from 3 to 4
	fmt.Println("fruitlis is", fruitList)
	fruitList = append(fruitList[:3]) // this will return values from to default to 3
	fmt.Println("fruitlis is", fruitList)

	highscores := make([]int, 4) // this will make slice of size 4 and if you try to add element beyond size given it will throw error
	highscores[0] = 333
	highscores[1] = 350
	highscores[2] = 456
	highscores[3] = 667
	fmt.Println(highscores)
	highscores = append(highscores, 567, 45, 789, 980, 345) //but by using append we can add elements to slices here
	fmt.Println(highscores)
	sort.Ints(highscores) //this will sort elements in increasing order
	fmt.Println(highscores)
	fmt.Println(sort.IntsAreSorted(highscores)) //this will return true or false based upon the elements are sorted or not

}
