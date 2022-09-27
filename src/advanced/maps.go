package main

import "fmt"

func main() {
	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"anfaas":  3434343,
		"testing": 343455,
	}
	fmt.Println(statePopulation["testing"])
	statePopulation["lolaland"] = 28238283
	delete(statePopulation, "anfaas")
	_, ok := statePopulation["lolaland"]
	fmt.Println(ok)
}
