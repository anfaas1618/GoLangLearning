package main

import (
	"fmt"
	"reflect"
)

type animal struct {
	Name   string `required max:"100"`
	Origin string
}

func main() {
	t := reflect.TypeOf(animal{})
	field, ok := t.FieldByName("Name")
	fmt.Println(field.Tag, ok)

}
