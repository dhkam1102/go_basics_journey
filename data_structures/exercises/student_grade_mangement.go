// Exercise: Student Grade Management System
// Problem Description:

// Create a student grade management system that allows users to interact with a database of student grades.
// The program should allow the user to perform basic operations such as adding a student,
// assigning grades for various subjects, updating grades, calculating average grades,
// and displaying all student grades. The goal is to use control structures,
// functions, arrays, slices, and maps effectively to manage the data and user input.

// Requirements:
// Initial Setup:

// Start with an empty list of students.
// Each student should have a name and a map of subjects with corresponding grades.
// Display a menu with options for the user: Add Student, Assign Grade,
// Update Grade, Calculate Average Grade, Display All Grades, and Exit.

// Control Structures:
// Use a loop to repeatedly show the menu until the user decides to exit.
// Handle the user's menu selection using control structures like if-else or switch-case.
// Functions:

// Add Student:

// Implement a function to add a new student to the list. The function should prompt the user to enter the student's name.
// Ensure no duplicate student names are allowed.
// Assign Grade:

// Implement a function to assign a grade to a student for a specific subject.
// If the subject doesn't exist for the student, add the subject and the grade to their record.
// Update Grade:

// Implement a function to update an existing grade for a specific subject of a student.
// Calculate Average Grade:

// Implement a function to calculate the average grade for a specific student across all subjects.
// Display All Grades:

// Implement a function to display all students and their grades in all subjects.
// Display Menu:

// Implement a function to display the menu and get the user's choice.
// Variables:

// Maintain a list of students, where each student is represented as a map entry
// with their name as the key and their subjects and grades as another map.
// Use variables to handle user input and other necessary data for transactions.
// User Interaction:

// The user should be able to repeatedly interact with the system until they choose to exit.
// The user can add students, assign or update grades, calculate averages, or display all grades based on their choice.
// The program should provide feedback on each operation, such as confirming a grade assignment or showing the calculated average.

package main

import (
	"fmt"
)

type student struct {
	name    string
	classes map[string]string
}

func addStudent(students *map[string]map[string]string) {
	fmt.Println("Enter student name: ")
	var name string
	fmt.Scan(&name)

	if _, exist := (*students)[name]; exist {
		fmt.Println("student already exists")
		return
	}
	(*students)[name] = make(map[string]string)
	fmt.Println("Student added successfully!!")
}

func assignGrade(students *map[string]map[string]string) {
	fmt.Println("Enter relevant info (name, subject, grade(A, B, C, D))")
	var name, subject, grade string
	fmt.Scanf("&s, &s, &s", &name, &subject, &grade)

	_, ok := (*students)[name][subject]
	if ok {
		fmt.Println("This student already has a grade for the given subject")
	} else {
		fmt.Println("Successfully added a grade for the given subject")
		(*students)[name][subject] = grade
	}
}

func updateGrade(students *map[string]map[string]string) {
	fmt.Println("Enter relevant info (name, subject, grade(A, B, C, D))")
	var name, subject, grade string
	fmt.Scanf("&s, &s, &s", &name, &subject, &grade)
	(*students)[name][subject] = grade
	fmt.Println("Successfully updated the grade")
}

func calculateAverage(students *map[string]map[string]string) float64 {
	fmt.Println("Enter student name: ")
	var name string
	fmt.Scan(&name)

	var total float64 = 0
	var num_classes float64 = 0
	for _, grade := range (*students)[name] {
		num_classes += 1
		switch grade {
		case "A":
			total += 4.0
		case "B":
			total += 3.0
		case "C":
			total += 2.0
		case "D":
			total += 1.0
		}
	}
	if num_classes == 0 {
		return 0
	}
	return total / num_classes
}

func displayAll(students *map[string]map[string]string) {

	for name, classes := range *students {
		var total float64 = 0
		var num_classes float64 = 0
		for _, grade := range classes {
			num_classes += 1
			switch grade {
			case "A":
				total += 4.0
			case "B":
				total += 3.0
			case "C":
				total += 2.0
			case "D":
				total += 1.0
			}
		}
		if num_classes == 0 {
			total = 0
		}
		total /= num_classes
		fmt.Printf("Student: %s, total: %.2f\n", name, total)
	}

}

func option() {
	fmt.Println("OPTIONS: \n" +
		"ADD STUDENT, ASSIGN GRADE, UPDATE GRADE\n" +
		"CALAC AVG GRADE, DISPLAY ALL GRADE, OPTIONS, EXIT")
}

func main() {

	var students map[string]map[string]string

	fmt.Println("Welcome to the Student System " +
		"You can manage your students with the following options:\n" +
		"ADD STUDENT, ASSIGN GRADE, UPDATE GRADE\n" +
		"CALAC AVG GRADE, DISPLAY ALL GRADE, OPTIONS, EXIT")

	var user_input string

	for {
		fmt.Println("Enter your option: ")
		fmt.Scan(&user_input)
		switch user_input {
		case "ADD STUDENT":
			addStudent(&students)
		case "ASSIGN GRADE":
			assignGrade(&students)
		case "UPDATE GRADE":
			addStudent(&students)
		case "CALAC AVG GRADE":
			calculateAverage(&students)
		case "DISPLAY ALL GRADE":
			displayAll(&students)
		case "OPTIONS":
			option()
		case "EXIT":
			return
		default:
			fmt.Println("invalid option")
		}
	}
}
