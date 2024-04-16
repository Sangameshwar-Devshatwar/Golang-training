package main

import "fmt"

var name string
var mobile string
var no int

var Tickets int = 55

var storeNames = make([]string, 0)
var userTickets int
var leftTickets int
var condition bool = true

func main() {

	// fmt.Println("1.register name and mobile number,
	// 			 2.Enter number of tickets you want to book ,
	// 			3.left tickets,
	// 			 4.List of all registered users")

	for condition {
		fmt.Println("1.registration\n2.tickets booking\n3.left tickets\n4.list of registered users\n5.for exit")
		fmt.Scan(&no)
		switch no {
		case 1:
			fmt.Println("Enter your name")
			fmt.Scan(&name)
			fmt.Println("Enter your mobile no")
			fmt.Scan(&mobile)

			storeNames = append(storeNames, name)

			fmt.Println("welcome to conference booking System mr.", name)
			break
		case 2:
			fmt.Println("enter the number of tickets you want Book")
			fmt.Scan(&userTickets)

			Tickets = Tickets - userTickets
			// if Tickets == 0 {
			// 	fmt.Println("all tickets sold")
			// 	break
			// }
			fmt.Println("your booking is successfully completed ")
			fmt.Println("remaining Number of tickets are", Tickets)
			break
		case 3:

			fmt.Println("Total number of tickets left are", Tickets)
			break
		case 4:
			fmt.Println("list of all registered users are", storeNames)
			break
		case 5:
			condition = false
			break
		}
	}
}
