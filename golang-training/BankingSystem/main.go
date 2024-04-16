// Banking System:

// Develop a CLI application to simulate a basic banking system.
// Create struct types for Account, Transaction, etc.
// Implement methods for opening an account, depositing funds, withdrawing funds, checking balance, etc.
// Utilize receivers for each method defined on the relevant types.
package main

import (
	"bankingsystem/account"
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {

	var balance float64
	var accnum = 33223414051
	var inputDate string
	var acctype string
	Bank := make(map[int]account.Account)
	reader := bufio.NewReader(os.Stdin)
	rand.Seed(time.Now().UnixNano())

	pin := rand.Intn(9000) + 1000

	for {
		var i int
		fmt.Print("1)create ac\n 2)set pin\n3)all ac details\n 4)deposit funds\n 5)withdraw funds\n 6)chk bal\n 7)transfer money\n 8)exit\n")
		fmt.Scan(&i)

		switch i {
		case 1:

			fmt.Println("_______Enter ac details_________")
			fmt.Print("Enter your name: ")

			input, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			fmt.Println("____Enter Date of Birth____")
			fmt.Scan(&inputDate)

			year, _ := strconv.Atoi(inputDate[:4])
			month, _ := strconv.Atoi(inputDate[5:7])
			day, _ := strconv.Atoi(inputDate[8:10])

			dateOnly := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)

			fmt.Println("Enter balance:")
			fmt.Scan(&balance)
			fmt.Println("Enter account type: [savings or current]")
			fmt.Scan(&acctype)
			// fmt.Println("Enter pin")
			// fmt.Scan(&pin)
			ac := account.Account{}
			ac.SetName(input)
			ac.SetBalance(balance)
			ac.GetACNum(accnum)
			ac.SetDob(dateOnly)
			ac.SetAccType(acctype)
			ac.GetPin(pin)
			Bank[accnum] = ac
			accnum++

		case 2: //set Pin
			var acnum int
			var userInput int
			var oldPin int
			fmt.Println("Enter account number")
			fmt.Scan(&acnum)
			if acnum == Bank[acnum].AcNum {
				fmt.Println("firstly Enter old pin to set new pin")
				fmt.Scan(&oldPin)
				if oldPin == Bank[acnum].Pin {
					fmt.Println("Enter your new pin to set: by changing old")
					fmt.Scan(&userInput)
					if userInput == Bank[accnum].Pin {
						fmt.Println("recent password set new one ")
					} else {
						account.SetPin(userInput, acnum, &Bank)

					}
				} else {
					fmt.Println("Enter valid account number")
				}
			}
		case 3:
			fmt.Println("All account details:")
			fmt.Println(Bank)
		case 4: //depositing funds
			var acc int
			var fund float64
			var InputPin int
			fmt.Println("Enter account number to deposit funds")
			fmt.Scan(&acc)
			fmt.Println("Enter correct pin to withdraw amount")
			fmt.Scan(&InputPin)
			if Bank[acc].AcNum == acc && Bank[acc].Pin == InputPin {
				fmt.Println("Enter amount to deposit")
				fmt.Scan(&fund)
				if fund > 0 {
					//value := bank[acc]
					if _, ok := Bank[acc]; ok {
						account.DepositAmount(fund, acc, &Bank)
					} else {
						fmt.Println("Invalid account")
					}
				} else {
					fmt.Println("Enter valid amount")
				}
			} else {
				fmt.Println("Enter valid acc number or pin")
			}

		case 5: //withdrawing funds
			var acc int
			var withdrawAm float64
			var InputPin int

			fmt.Println("Enter acc number to withdraw fund")
			fmt.Scan(&acc)
			fmt.Println("Enter correct pin to withdraw amount")
			fmt.Scan(&InputPin)
			if Bank[acc].AcNum == acc && Bank[acc].Pin == InputPin {
				fmt.Println("Enter amount to withdraw")
				fmt.Scan(&withdrawAm)
				if _, ok := Bank[acc]; ok {
					account.WithdrawAmount(withdrawAm, &Bank, acc)
				}
			} else {
				fmt.Println("Invalid Account no or Invalid pin") //it is useful to be safe from hackers by confusing them
			}
		case 6: //viewing Balance
			var acc int
			var InputPin int
			fmt.Println("Enter acc number to check balance")
			fmt.Scan(&acc)
			fmt.Println("Enter correct pin to check Balance")
			fmt.Scan(&InputPin)
			if Bank[acc].AcNum == acc && Bank[acc].Pin == InputPin {

				amount := account.CheckBalance(acc, &Bank)
				fmt.Printf("Balance is %f", amount)
			} else {
				fmt.Println("Invalid account number or Invalid pin ")
			}
		case 7: //transferring money to another account
			var acc1 int
			var acc2 int
			var InputPin1 int
			var amount float64

			fmt.Println("Enter acc number of yours to transfer funds")
			fmt.Scan(&acc1)
			fmt.Println("Enter pin:: ")
			fmt.Scan(&InputPin1)

			if Bank[acc1].AcNum == acc1 && Bank[acc1].Pin == InputPin1 {
				fmt.Println("Enter acc number of the other person where you want to transfer funds")
				fmt.Scan(&acc2)
				if acc2 == Bank[acc2].AcNum {
					fmt.Println("Enter amount to transfer")
					fmt.Scan(&amount)
					money := account.CheckBalance(acc1, &Bank)
					if amount > money {
						fmt.Println("Insufficient Funds")

					} else {
						account.WithdrawAmount(amount, &Bank, acc1)
						account.DepositAmount(amount, acc2, &Bank)
					}
				}

			} else {
				fmt.Println("Enter valid acc number to transfer funds")
			}

		case 8:
			return
		}
	}
}
