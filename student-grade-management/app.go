package main

import "fmt"

var students = make(map[string]float64)
var studentName string
var studentGrade float64

func main() {
	fmt.Println("Hello, World!")
	var choice int
	for choice != 5 {
		ShowMenu()
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Add student")
			AddStudent()
		case 2:
			fmt.Println("Update Grade")
			UpdateGrade()
		case 3:
			fmt.Println("Show All Grade")
			ShowAllGrade()
		case 4:
			fmt.Println("Show Overall Grade")
			ShowOverallGrade()
		case 5:
			fmt.Println("Exiting..")
		default:
			fmt.Println("Invalid choice")
		}
	}			
}


func ShowMenu() {
	fmt.Println("\n\n\n1. Add student")
	fmt.Println("2. Update Grade")
	fmt.Println("3. Show All Grade")	
	fmt.Println("4. Show Averall Grade")
	fmt.Println("5. Exit")
}

func AddStudent(){
	var name string
	var grade float64

	fmt.Println("Enter student name:")
	fmt.Scan(&name)

	fmt.Println("Enter student grade:")
	fmt.Scan(&grade)

	if _, ok := students[name]; ok {
		fmt.Println("Student already exists")
		return
	}else{
		students[name] = grade
		ClearScren()
		fmt.Printf("Student %v added successfully\n", name)
	}
}

func UpdateGrade(){
	var name string
	var grade float64

	fmt.Println("Enter student name:")
	fmt.Scan(&name)

	fmt.Println("Enter student grade:")
	fmt.Scan(&grade)

	if _, ok := students[name]; ok {
		students[name] = grade
		ClearScren()
		fmt.Printf("Student %v grade updated successfully\n", name)
	}else{
		ClearScren()
		fmt.Println("Student does not exists")
	}
}

func ShowAllGrade(){
	ClearScren()
	// Format table for students and grades
	fmt.Println("Name\tGrade")
	fmt.Println("----\t-----")
	for name, grade := range students {
		fmt.Printf("%v\t%v\n", name, grade)
	}
}

func ShowOverallGrade(){
	var total float64
	var count int

	for _, grade := range students {
		total += grade
		count++
	}
	average := total / float64(count)
	ClearScren()
	fmt.Printf("Overall grade: %.2f\n", average)
}

func ClearScren(){
	fmt.Print("\033[H\033[2J")
}
