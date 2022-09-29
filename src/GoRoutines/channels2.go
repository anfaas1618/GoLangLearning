package main

import "fmt"

func foo(someValue int, c chan int) {
	c <- someValue * 5
}
func main() {
	fooVal := make(chan int)
	go foo(2, fooVal)
	go foo(3, fooVal)
	go foo(4, fooVal)

	v1 := <-fooVal
	v2 := <-fooVal
	v3 := <-fooVal
	v4 := <-fooVal

	fmt.Println(v1, v2, v3, v4)
}
