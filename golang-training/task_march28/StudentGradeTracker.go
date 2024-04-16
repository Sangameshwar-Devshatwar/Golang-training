/*"Q. Student Grade Tracker:
Build an application to track student grades for multiple subjects.
Allow users to add grades, calculate average grades, and determine the overall grade."*/

package main

import "fmt"

var condition bool = true
var nos int

//gradesMap:=make(map[string]string)

func main() {
	var n string
	var roll int
	var no int
	var subEnglish int
	var subPhysics int
	var subChem int
	var details []Student
	for condition {

		fmt.Println("1.enter student details\n2.student details\n3.watch grades\n4.for exit")
		fmt.Scan(&nos)
		switch nos {
		case 1:

			fmt.Println("Enter name of Student")
			fmt.Scan(&n)

			fmt.Println("Enter roll no of Student")
			fmt.Scan(&roll)

			fmt.Println("Enter std of Student")
			fmt.Scan(&no)

			fmt.Printf("Enter marks of english of student name:%s\n", n)
			fmt.Scan(&subEnglish)
			fmt.Printf("Enter marks of physics of student name:%s\n", n)
			fmt.Scan(&subPhysics)

			fmt.Printf("Enter marks of chemistry of student name:%s\n", n)
			fmt.Scan(&subChem)
			fmt.Println("----Details added successfully-------")
		case 2:
			subjectDetails := Subjects{
				English:   subEnglish,
				Physics:   subPhysics,
				Chemistry: subChem}
			studentDetails := Student{
				Name:     n,
				RollNo:   roll,
				Std:      no,
				Subjects: subjectDetails}
			details = append(details, studentDetails)
			fmt.Println(details)

		case 3:
			if subEnglish < 60 {

				fmt.Println("grade for english is: C grade \n")
			} else if subEnglish > 80 {
				fmt.Println("grade for english is: A grade \n")

			} else {
				fmt.Println("grade for english is: B grade \n")
			}
			if subPhysics < 60 {

				fmt.Println("grade for physics is: C grade\n")
			} else if subPhysics > 80 {
				fmt.Println("grade for physicsis: A grade")

			} else {
				fmt.Println("grade for physics is: B grade")
			}
			if subChem < 60 {

				fmt.Println("grade for chemistry is: C grade")
			} else if subChem > 80 {
				fmt.Println("grade for chemistry is: A grade")

			} else {
				fmt.Println("grade for chemistry is: B grade")
			}
			Total := subEnglish + subPhysics + subChem
			fmt.Printf("total subjects marks are:%d\n", Total)
			per := Total / 3
			fmt.Printf("total percentage is:%d\n", per)

			if per < 60 {

				fmt.Println("overall grade is: C grade")
			} else if per > 80 {
				fmt.Println("overall grade is: A grade")

			} else {
				fmt.Println("overall grade is: B grade")
			}

		case 4:
			condition = false

		}
	}
}

type Student struct { //strcture for student data
	Name     string
	RollNo   int
	Std      int
	Subjects Subjects
}
type Subjects struct {
	English   int
	Physics   int
	Chemistry int
}
