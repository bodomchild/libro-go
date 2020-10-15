package ej4_3_4_7

import "unicode"

// Reverse is the exercise 4.3
func Reverse(p *[]int) {
	for i, j := 0, len(*p)-1; i < j; i, j = i+1, j-1 {
		(*p)[i], (*p)[j] = (*p)[j], (*p)[i]
	}
}

// Rotate is the excersise 4.4
func Rotate(s []int, i int) {
	s2, s3 := s[:i], s[i:]
	Reverse(&s2)
	Reverse(&s3)
	Reverse(&s)
}

// RemoveAdjacentDuplicates is the excercise 4.5
func RemoveAdjacentDuplicates(strings []string) []string {
	i := 1
	for _, s := range strings {
		if s != strings[i-1] {
			strings[i] = s
			i++
		}
	}
	return strings[:i]
}

// SquashAdjacentSpace is the excercise 4.6
func SquashAdjacentSpace(s []byte) []byte {
	i := 1
	for _, b := range s {
		if !unicode.IsSpace(rune(b)) || !unicode.IsSpace(rune(s[i-1])) {
			s[i] = b
			i++
		}
	}
	return s[:i]
}

// SquashAdjacentSpace is the excercise 4.7
func ReverseChars(b []byte) {
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
}
