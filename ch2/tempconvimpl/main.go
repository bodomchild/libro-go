package main

import (
	"awesomeProject2/ch2/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i, arg := range os.Args[1:] {
		t, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "cf: %v\n", err)
			os.Exit(1)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)
		fmt.Printf("%d-\t%s = %s, %s = %s\n", i, f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}
