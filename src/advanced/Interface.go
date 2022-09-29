package main

import "fmt"

func main() {
	w := ConsoleWriter{}
	_, err := w.Write([]byte("hello world"))
	if err != nil {
		return
	}
}

type Writer interface {
	write([]byte) (int, error)
}

type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err

}
