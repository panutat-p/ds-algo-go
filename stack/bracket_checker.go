package stack

import (
	"container/list"
	"fmt"
)

// use stack to collect left bracket and pop to test right bracket

func IsBalancedBrackets(str string) bool {
	leftBrackets := map[string]string{"(": ")", "[": "]", "{": "}", "<": ">"}
	rightBrackets := map[string]bool{")": true, "]": true, "}": true, ">": true}

	s := Stack{list.New()}
	for _, v := range str {
		b := string(v)
		if right, ok := leftBrackets[b]; ok {
			s.Push(right)
			s.Print()
		} else if rightBrackets[b] {
			if s.IsEmpty() || s.Pop() != b { // test right bracket
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
