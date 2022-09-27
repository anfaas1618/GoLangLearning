package main

import "fmt"

func main() {
	name := "hello world"
	var _ string = "testingstrings"
	fmt.Println(name)

	for _, v := range name {
		fmt.Println(string(v))
	}
}
