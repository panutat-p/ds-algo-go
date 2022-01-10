package stack

import (
	"fmt"
)

// IsBalancedBrackets
// when found left bracket, look for right bracket and store to stack
func IsBalancedBrackets(str string) bool {
	leftBrackets := map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}
	rightBrackets := map[string]bool{")": true, "]": true, "}": true, ">": true}

	s := New()
	for _, v := range str {
		b := string(v)
		if right, ok := leftBrackets[b]; ok { // it is left bracket
			s.Push(right)
			s.Print()
		} else if rightBrackets[b] { // it is right bracket
			if s.IsEmpty() || s.Pop() != b {
				fmt.Println("found mismatch char:", b)
				return false
			}
		} else { // ignore other characters
			fmt.Println("ignore", b)
			continue
		}
	}
	return s.IsEmpty()
}
