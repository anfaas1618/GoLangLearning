package main

import "fmt"

func main() {
	grades := [...]string{"ola", "amo", "samo", "kola", "sola", "torola"}
	var student [3]string
	grades[2] = "hello"
	student[0] = "lisa"
	fmt.Printf("%v , %v\n", grades, student)

	///slice

	gradeSlice := []int{1, 2, 3}
	gradeSlice = append(gradeSlice, 23)

	for _, v := range grades {
		fmt.Println(v)
	}

}
