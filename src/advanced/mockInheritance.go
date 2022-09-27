package main

import "fmt"

type Animal struct {
	name string
	age  int
}
type bird struct {
	Animal
	wings  int
	canFly bool
}

// WHEN TO USE -
func main() {
	dodo := bird{
		Animal: Animal{
			name: "dodo",
			age:  22,
		},
		wings:  2,
		canFly: false,
	}
	fmt.Println(dodo)
}
