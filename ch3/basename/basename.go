package basename

import (
	"bytes"
	"fmt"
	"strings"
)

func Basename(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func Basename2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot >= 0 {
		s = s[:dot]
	}
	return s
}

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

// Ej 3.10, copiado de internet
func CommaNonRecursive(s string) string {
	var buf bytes.Buffer
	pre := len(s) % 3

	if pre == 0 {
		pre = 3
	}

	buf.WriteString(s[:pre])

	for i := pre; i < len(s); i += 3 {
		_, _ = fmt.Fprintf(&buf, ",%s", s[i:i+3])
	}

	return buf.String()
}

// Ej 3.11
func EnhancedComma(s string) string {
	var buf bytes.Buffer
	var sign, decimalPart string
	pre := len(s) % 3

	if strings.HasPrefix(s, "-") {
		sign = "-"
	}
	if strings.Contains(s, ".") {
		decimalPart = strings.Split(s, ".")[1]
	}

	if pre == 0 {
		pre = 3
	}

	s = strings.TrimPrefix(strings.TrimSuffix(s, "."+decimalPart), sign)
	buf.WriteString(s[:pre])

	for i := pre; i < len(s); i += 3 {
		_, _ = fmt.Fprintf(&buf, ",%s", s[i:i+3])
	}

	s = buf.String()

	if sign == "-" {
		s = sign + s
	}
	if decimalPart != "" {
		s = s + "." + decimalPart
	}

	return s
}
