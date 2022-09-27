package main

import (
	"fmt"
)

func main() {
	a := 34
	var b *int = &a
	*b++
	fmt.Println(a, *b)

	//array
	x := [3]int{1, 2, 3}
	y := &x[0]
	z := &x[1]
	fmt.Println(x, &y, z)
	//
	var ms *myStruct
	ms = new(myStruct)
	ms.foo = 34
	fmt.Println(ms.foo)

}

type myStruct struct {
	foo int
}
