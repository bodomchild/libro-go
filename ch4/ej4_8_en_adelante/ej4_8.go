package ej4_8_en_adelante

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

// CharCategoryCount is the excercise 4.8
func CharCategoryCount(input *bufio.Reader) {
	counts := make(map[string]int)
	invalid := 0

	for {
		r, n, err := input.ReadRune()
		if err == io.EOF {
			break
		}
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		switch {
		case unicode.IsLetter(r):
			counts["letter"]++
		case unicode.IsNumber(r):
			counts["number"]++
		case unicode.IsSpace(r):
			counts["space"]++
		case unicode.IsControl(r):
			counts["control"]++
		}
	}

	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

// Wordfreq is the excercise 4.9
func Wordfreq(input io.Reader) {
	count := make(map[string]int)
	scan := bufio.NewScanner(input)
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		p := scan.Text()
		count[p]++
	}

	for word, qt := range count {
		fmt.Printf("%-30s %d\n", word, qt)
	}
}
