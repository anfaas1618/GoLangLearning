package main

import "fmt"

func Calculate(x int) (result int) {

	result = x + 2
	return
}
func main() {
	fmt.Println("go testing tutoral")
	result := Calculate(2)
	fmt.Println(result)
}
