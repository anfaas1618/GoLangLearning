package main

import "fmt"

func main() {
	slice := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	a := slice[:]
	b := slice[3:len(slice)]
	c := slice[:4]
	d := slice[2:7]
	slice[3] = 555
	fmt.Println(slice)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	//make
	m := make([]int, 100, 555)
	m = append(m, 55)
	fmt.Println(m)
	fmt.Println(len(m))
	fmt.Println(cap(m))
}
