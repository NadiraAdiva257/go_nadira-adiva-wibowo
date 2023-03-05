package main

import (
	"fmt"
)

func main() {
	listStudent := Student{
		name:  []string{"Rizky", "Kobar", "Ismail", "Umam", "Yopan"},
		score: []int{80, 75, 70, 75, 60},
	}

	for i := 0; i < len(listStudent.name); i++ {
		fmt.Println("Input", i+1, "Student's Name", listStudent.name[i])
		fmt.Println("Input", i+1, "Student's Score", listStudent.score[i])
	}

	fmt.Println()

	fmt.Println("Average Score:", listStudent.Average())

	min, max, nameMin, nameMax := listStudent.MinMaxScore()
	fmt.Println("Min Score of Students:", nameMin, "-", min)
	fmt.Println("Max Score of Students:", nameMax, "-", max)
}

type Student struct {
	name  []string
	score []int
}

func (student Student) Average() int {
	sum := 0

	for _, value := range student.score {
		sum += value
	}

	return sum / len(student.score)
}

func (student Student) MinMaxScore() (min, max int, nameMin, nameMax string) {
	min = student.score[0]
	max = student.score[0]
	nameMin = student.name[0]
	nameMax = student.name[0]

	for i := 0; i < len(student.score); i++ {
		if student.score[i] < min {
			min = student.score[i]
			nameMin = student.name[i]
		}

		if student.score[i] > max {
			max = student.score[i]
			nameMax = student.name[i]
		}
	}

	return
}
