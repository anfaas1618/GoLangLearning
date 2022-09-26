package main

import "fmt"

func main() {
	if 7%2 == 0 {
		println("7 is even")
	} else {
		println("7 is odd")
	}
	if 7 > 2 {
		println("hello")
	}
	if num := 9; num < 0 {
		println(num)
	} else if num > 4 {
		println("yoo")
	}
	var a [5]int
	fmt.Println("emp:", a)
	a[4] = 100
	fmt.Println("set:", a)
	fmt.Println("get:", a[4])

}
