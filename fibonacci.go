package main

import "fmt"

func main() {
	f := Fibonacci()
	fmt.Print(f())
}

func Fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}
