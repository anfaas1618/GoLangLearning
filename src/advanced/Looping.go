package main

import "fmt"

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
