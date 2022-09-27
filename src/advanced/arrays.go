package main

import "fmt"

func main() {
	grades := [...]string{"ola", "amo", "samo"}
	var student [3]string
	grades[2] = "hello"
	student[0] = "lisa"
	fmt.Printf("%v , %v", grades, student)

	///slice

	gradeSlice := []int{1, 2, 3}
	gradeSlice = append(gradeSlice, 23)

}
