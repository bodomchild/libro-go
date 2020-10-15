package main

import (
	"awesomeProject2/ch3/basename"
	"awesomeProject2/ch3/constants"
	"awesomeProject2/ch3/printints"
	"fmt"
)

func main() {
	fmt.Println(printints.IntsToString([]int{1, 2, 3, 4, 5}))
	fmt.Println(basename.Comma("1234567890"))
	fmt.Println(basename.CommaNonRecursive("1234567890"))
	fmt.Println(basename.EnhancedComma("-1234567890.1234"))

	constants.Mostrar2()
	constants.Mostrar10()
}
