package main

import (
	"fmt"
)

func main() {
	i := 1
	switch i {
	case 1:
		fmt.Println("first")
		fallthrough
	case 2:
		fmt.Println("second")
	default:
		fmt.Println("enter a valid number")

	}
	//switch time.Now().Weekday() - 1 {
	//case time.Saturday, time.Sunday:
	//	fmt.Println("weekend")
	//case time.Monday:
	//	fmt.Println("its monday")
	//default:
	//	fmt.Println("other days")
	//}
}
