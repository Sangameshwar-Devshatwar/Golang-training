package main

import "fmt"

type Product struct {
	Product_name  string
	Product_Price int
}
type Product_Details struct {
	Collection map[int]Product
}

func main() {
	var PCBrandid int 

	var PCname string
	var PCprice int
	//var condition bool = true
	details := Product_Details{
		Collection: make(map[int]Product),
	}
	for {
		var val int

		fmt.Println("(1)Add product details\n 2)show all products list\n3)Remove product\n4)update name of product\n5)Update by price\n 6)EXIT")
		fmt.Scan(&val)
		switch val {

		case 1:

			fmt.Println("enter pc id")
			fmt.Scan(&PCBrandid)

			fmt.Println("Enter PCname ")
			fmt.Scan(&PCname)
			fmt.Println("enter pc price")
			fmt.Scan(&PCprice)
			All_Details := Product{
				Product_name:  PCname,
				Product_Price: PCprice,
			}
			details.Collection[PCBrandid] = All_Details
		case 2:

			fmt.Println(details.Collection)
		case 3:

			deleteByKey(&details)
		case 4:

			updateByName(&details)

		case 5:
			updateByPrice(&details)
		case 6:
			return

		}
	}
}

func deleteByKey(details *Product_Details) {
	var pid int
	fmt.Println("Enter Product id to delete particular product of that brand")
	fmt.Scan(&pid)
	delete(details.Collection, pid)

}
func updateByName(details *Product_Details) {
	var name string
	var id int

	fmt.Println("eneter id where you want to update name")
	fmt.Scan(&id)
	fmt.Println("Enter Product name that you want to update ")
	fmt.Scan(&name)
	if p, ok := details.Collection[id]; ok {
		p.Product_name = name
		details.Collection[id] = p
	}
	fmt.Println(" name updated successfully")

}
func updateByPrice(details *Product_Details) {
	var id int
	var price int

	fmt.Println("enter id where you want to update price")
	fmt.Scan(&id)
	fmt.Println("Enter Product price that you want to update ")
	fmt.Scan(&price)
	if p, ok := details.Collection[id]; ok {
		p.Product_Price = price
		details.Collection[id] = p
	}
	fmt.Println(" Price updated successfully")

}
