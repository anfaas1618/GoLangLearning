package main

import "fmt"

const (
	isAdmin = 1 << iota
	isHead
	isFinance
	isTeacher
	isAVP
)

func main() {
	var roles byte = isAdmin | isHead | isAVP
	fmt.Printf("%b", roles)
}
