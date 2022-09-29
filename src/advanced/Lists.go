package main

import (
	"container/list"
	"fmt"
)

func main() {
	l := list.New()
	fmt.Println(l)
	l.PushFront(34)
	l.PushFront(55)
	l.PushBack(77)
	fmt.Println(l.Front())
	fmt.Println(l.Back())
	l.Remove(l.Front().Next())
	//
	fmt.Println("print")
	for i := l.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}
}
