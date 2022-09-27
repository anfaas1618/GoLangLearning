package main

import "fmt"

func main() {
	statePopulation := make(map[string]int)
	statePopulation = map[string]int{
		"anfaas":  3434343,
		"testing": 343455,
		"ohio":    232954,
	}
	fmt.Println(statePopulation["testing"])
	statePopulation["lalaland"] = 28238283
	delete(statePopulation, "anfaas")
	_, ok := statePopulation["lalaland"]
	fmt.Println(ok)

	fmt.Println("**********************")
	for k, v := range statePopulation {
		fmt.Println("the key is ", k)
		fmt.Println("the value is ", v)

	}
}
