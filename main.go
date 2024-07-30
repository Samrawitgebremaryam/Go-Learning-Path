package main

import (
	"fmt"
)

var grade map[string]int
var Name string
var subjectNumber int

func main() {
	input()
	average := gradeCalculator(grade, subjectNumber)
	display(subjectNumber, Name, average)
}

func input() {
	grade = make(map[string]int)
	var Subject string
	var Value int

	fmt.Print("Hello There, Enter your name please: ")
	fmt.Scan(&Name)

	fmt.Print("Enter how many subjects you take: ")
	fmt.Scan(&subjectNumber)

	for i := 0; i < subjectNumber; i++ {
		fmt.Print("Please enter the name of the subject: ")
		fmt.Scan(&Subject)

		for {
			fmt.Print("Please enter your grade (0-100): ")
			fmt.Scan(&Value)
			if Value >= 0 && Value <= 100 {
				break
			}
			fmt.Println("Invalid grade! Please enter a value between 0 and 100.")
		}

		grade[Subject] = Value
	}
}

func gradeCalculator(grade map[string]int, subjectNumber int) float64 {
	total := 0
	for _, val := range grade {
		total += val
	}

	if subjectNumber == 0 {
		return 0
	}
	average := float64(total) / float64(subjectNumber)
	return average
}

func display(subjectNumber int, Name string, average float64) {
	fmt.Println("\t\t\t########################################")
	fmt.Printf("\t\t\tName of student: %s\n", Name)
	fmt.Printf("\t\t\tNumber of subjects: %d\n", subjectNumber)
	fmt.Println("\t\t\tGrades entered:")

	for subject, grade := range grade {
		fmt.Printf("\t\t\t%s: %d\n", subject, grade)
	}

	fmt.Printf("\t\t\tYour average is: %.2f\n", average)
	fmt.Println("\t\t\t########################################")
}
