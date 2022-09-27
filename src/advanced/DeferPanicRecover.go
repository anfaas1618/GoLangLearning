package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

/*
summary
defer runs in lifo order and only executes when full function ends
*/
func main() {
	fmt.Println("start")
	defer fmt.Println("middle")
	fmt.Println("end")
	defer fmt.Println("middle2")

	res, err := http.Get("https://www.google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(robots))

}
