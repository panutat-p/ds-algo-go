package stack

import (
	"fmt"
	"github.com/panutat-p/fiset-complete-ds-go/queue"
	"strings"
	"unicode"
)

// IsPalindrome
// use only stack
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

// IsPalindromeQueueStack
// use both queue and stack
// this code is cleaner
func IsPalindromeQueueStack(str string) bool {
	s := New()
	q := queue.New()

	for _, v := range str {
		if unicode.IsLetter(v) {
			ele := strings.ToLower(string(v))
			s.Push(ele)
			q.Enqueue(ele)
		}
	}

	for !s.IsEmpty() && !q.IsEmpty() {
		if s.Pop() != q.Dequeue() {
			return false
		}
	}
	return true
}
