package main

import "fmt"

/*
Summary

	 *For statement
		1.simple

loop

		a.for initializer; test; incrementor {}
		b.for test {}
		c.for {}
	2.exiting early
		a.break
		b.continue
		c.labels
	3.looping over collection
		arrays,slices,maps,string,channels
		for k,b := range collection{ }
*/
func main() {

	//simple loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	//exit early
	for i := 0; i < 19; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}
	//label
Loop:
	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			println(i * j)
			if i*j > 66 {
				break Loop
			}
		}
	}
	println("this prints")
	//loop through collections
	s := []string{
		"hello",
		"ola",
		"amigos",
	}
	for _, v := range s {
		fmt.Println(v)
	}

}
