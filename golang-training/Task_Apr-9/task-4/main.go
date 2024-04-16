// Task 4: Struct Usage - Define a struct named Employee with fields Name, Age, and Salary.
//  Write a program that creates an array of Employee instances
// and prints the details of the employee with the highest salary.

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Employee struct {
	Id     int
	Name   string
	Age    int
	Salary float64
}

func (e *Employee) SetId(id int) {
	e.Id = id
}
func (e *Employee) SetName(name string) {
	e.Name = name
}
func (e *Employee) SetAge(age int) {
	e.Age = age
}
func (e *Employee) SetSalary(salary float64) {
	e.Salary = salary
}

func HighSalariedEmployee(EmployeeArr *[]Employee, sal float64) {
	for _, val := range *EmployeeArr {
		if val.Salary == sal {
			fmt.Println(val)
		}
	}
}

func main() {
	var id int
	reader := bufio.NewReader(os.Stdin) //used to add name that we can add full name using bufio,soo we have to declare it firstly to use
	var age int
	var sal float64

	var EmployeeArr []Employee
	var Salaries []float64

	for {
		var i int

		fmt.Println("Enter below option you want to perform operation\n")
		fmt.Println("1).add emp\n2).show all emp.3)show Emploee with highest salary  \n")
		fmt.Scan(&i)
		switch i {
		case 1:

			fmt.Println("Enter employee details")
			fmt.Println("********------------------------------------------********")
			fmt.Println("Enter id of Employee")
			fmt.Scan(&id)
			fmt.Println("Enter name of Employee")
			name, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("Error reading input:", err)
				return
			}
			fmt.Println("Enter Age of Employee")
			fmt.Scan(&age)
			fmt.Println("Enter Salary of Employee")
			fmt.Scan(&sal)

			details := Employee{}
			details.SetId(id)
			details.SetName(name)
			details.SetAge(age)
			details.SetSalary(sal)

			EmployeeArr = append(EmployeeArr, details)

		case 2:
			fmt.Println("-----All Employee Details----- ")
			fmt.Println(EmployeeArr)

		case 3:
			fmt.Println("--Employee with highest salary-- ")
			for _, val := range EmployeeArr {
				Salaries = append(Salaries, val.Salary)
			}
			sort.Float64s(Salaries)
			HighSalariedEmployee(&EmployeeArr, Salaries[len(Salaries)-1])

		}
	}
}
