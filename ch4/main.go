package main

import (
	"awesomeProject2/ch4/ej4_8_en_adelante"
	"fmt"
	"log"
	"os"
)

//
//import (
//	"awesomeProject2/ch4/ej4_1"
//	. "awesomeProject2/ch4/ej4_2"
//	"awesomeProject2/ch4/ej4_3-4_7"
//	"awesomeProject2/ch4/ej4_8_en_adelante"
//	"bufio"
//	"flag"
//	"fmt"
//	"os"
//)
//
//var input = flag.String("i", "", "string to encode")
//var sha = flag.Uint("s", 256, "type of SHA algorithm to use")
//
//func main() {
//	fmt.Println(ej4_1.DifferentBitsSHA256([]byte("holis"), []byte("hola")))
//
//	flag.Parse()
//
//	switch *sha {
//	case 256:
//		fmt.Println(SHA256Enc([]byte(*input)))
//	case 384:
//		fmt.Println(SHA384Enc([]byte(*input)))
//	case 512:
//		fmt.Println(SHA512Enc([]byte(*input)))
//	default:
//		fmt.Println("invalid argument for flag -s")
//	}
//
//	s := []int{1, 2, 3, 4, 5}
//	fmt.Printf("%d\n", s)
//	ej4_3_4_7.Reverse(&s)
//	fmt.Printf("%d\n", s)
//	ej4_3_4_7.Rotate(s, 2)
//	fmt.Printf("%d\n", s)
//
//	strings := []string{"Hola", "como", "como", "estas", "estas", "capo", "como", "andas"}
//	fmt.Println(ej4_3_4_7.RemoveAdjacentDuplicates(strings))
//
//	bytes := []byte("  hola \fcomo 	andas\n\ncapo")
//	fmt.Println(len(bytes))
//	bytes = ej4_3_4_7.SquashAdjacentSpace(bytes)
//	fmt.Printf("%s\n", bytes)
//	fmt.Println(len(bytes))
//
//	asd := []byte("Fernando")
//	fmt.Printf("%s\n", asd)
//	ej4_3_4_7.ReverseChars(asd)
//	fmt.Printf("%s\n", asd)
//
//	file, err := os.Open("prueba.txt")
//	if err != nil {
//		_,_ = fmt.Fprintf(os.Stderr, "main: %v\n", err)
//	}
//	defer func() {
//		_ = file.Close()
//	}()
//
//	reader := bufio.NewReader(file)
//	ej4_8_en_adelante.CharCategoryCount(reader)
//
//	ej4_8_en_adelante.Wordfreq(os.Stdin)
//}

func main() {
	result, err := ej4_8_en_adelante.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}
