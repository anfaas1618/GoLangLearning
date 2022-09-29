package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	go counts("really", c)
	for {
		msg, open := <-c

		if !open {

		}
		fmt.Println(msg)
	}
}

func counts(s string, c chan string) {
	for i := 0; i < 10; i++ {
		c <- s
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
}
