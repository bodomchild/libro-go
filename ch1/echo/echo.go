package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//args := os.Args[1:] // Original
	args := os.Args  // Ej 1.1
	switch args[1] { // args[0] para el original
	case "-e=1":
		echo1(args)
	case "-e=2":
		echo2(args)
	case "-e=3":
		echo3(args)
	case "-e=ej1-2":
		echo4(args) // Ej 1.2
	default:
		fmt.Println("Invalid first argument")
	}
}

func echo1(args []string) {
	var s, sep string
	for i := 0; i < len(args); i++ {
		s += sep + args[i]
		sep = " "
	}
	fmt.Println(s)
}

func echo2(args []string) {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func echo3(args []string) {
	fmt.Println(strings.Join(args, " "))
}

func echo4(args []string) {
	for index, arg := range args {
		fmt.Println(strings.Join([]string{strconv.Itoa(index), arg}, " "))
	}
}

// TODO ejercicio 1.3
