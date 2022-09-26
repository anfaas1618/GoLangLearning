package main

import (
	"database/sql/driver"
	"fmt"
)

func main() {
	var x int = 23
	var y = 23.2323
	var m = true
	var s = "hello9o" + "world"
	var n driver.Null
	fmt.Println(x, y, m, s, n)
}
