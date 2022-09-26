package main

import (
	"fmt"
	"math"
)

const x = "constnt"

func main() {
	const n = 2000000
	const m = 23 / float64(n)
	fmt.Println(x)
	fmt.Println(m)
	fmt.Println(math.Sin(m))
}
