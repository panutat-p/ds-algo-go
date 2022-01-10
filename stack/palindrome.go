package stack

import (
	"fmt"
	"strings"
	"unicode"
)

func IsPalindrome(str string) bool {
	var sb1 strings.Builder
	s := New()              // stack
	for _, v := range str { // v is type rune
		if unicode.IsLetter(v) {
			sb1.WriteString(strings.ToLower(string(v)))
			s.Push(strings.ToLower(string(v)))
		}
	}
	fmt.Println("sb1:", sb1.String())

	var sb2 strings.Builder
	for !s.IsEmpty() {
		if c, ok := s.Pop().(string); ok {
			sb2.WriteString(c)
		} else {
			panic("stack contains non string value")
		}
	}
	fmt.Println("sb2:", sb2.String())
	return sb1.String() == sb2.String()
}
