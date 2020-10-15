package main

import "fmt"

func main() {
	f := fibo()
	for i := 0; i < 90; i++ {
		fmt.Println(f())
	}
}

func fibo() func() int {
	a, b := 0, 1
	return func() int {
		c := a + b
		a, b = b, c
		return a
	}
}
