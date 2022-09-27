package main

import "fmt"

func main() {
	greet := greeter{
		greeting: "hello",
		name:     "hjello",
	}
	greet.greet()
	greet.greetChange()
	func() {
		fmt.Println("invoked")
	}()
	fmt.Println("hello world")
	fmt.Println(hello(99, " lol"))
	s := sum(1, 2, 3, 4, 5, 6, 7, 8, 8)
	fmt.Println(s)

	ans, err := divide(4, 4)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ans)
}

func hello(hello int, testing string) string {
	return "hello " + string(rune(hello)) + testing
}
func greeting(hello, testing, asd string) string {
	return hello + testing + asd
}

// variadic
func sum(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		fmt.Println(v)
		result += v
	}
	return
}
func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil

}

// method
type greeter struct {
	greeting string
	name     string
}

func (g greeter) greet() {

	fmt.Println(g.greeting, g.name)
	g.greeting = ""
}
func (g *greeter) greetChange() {

	fmt.Println(g.greeting, g.name)
	g.greeting = ""
}
