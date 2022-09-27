package main

import "fmt"

func main() {
	//	var n bool = true

	//every variable is initialized as 0
	//	fmt.Printf("%v %T ", n, n)

	a := 66.0
	b := 23.0
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	//	fmt.Println(a % b)
	fmt.Println(a + b)
	s := "this is a string"
	var c = []byte(s)
	fmt.Println(c)
}
