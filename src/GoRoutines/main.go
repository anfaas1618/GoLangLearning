package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		count("Fish")

		wg.Done()
	}()
	count2("sheep")
	wg.Wait()
}

func count(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "-", i)
		time.Sleep(700 * time.Millisecond)
	}
}

func count2(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "-", i)
		time.Sleep(500 * time.Millisecond)
	}
}
