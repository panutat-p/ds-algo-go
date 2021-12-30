package stack

import (
	"fmt"
	"testing"
)

func TestStackStoreString(t *testing.T) {
	s := New()
	s.Push("1")
	s.Push("2")
	s.Push("3")
	s.Print()
	fmt.Printf("%T\n", s.Pop())
}

func TestLoopThroughString(t *testing.T) {
	str := "{}[]{}[]()"
	for i := 0; i < len(str); i += 1 {
		fmt.Println(string(str[i]))
	}
}

func TestIsBalancedBrackets(t *testing.T) {
	s := "{}[]&{}[2]()"
	rslt := IsBalancedBrackets(s)

	if !rslt {
		t.Errorf("got %v, but expect %v", rslt, true)
	}
}

func TestIsBalancedBrackets_nested(t *testing.T) {
	s := "{{1{[()]}}2}"
	rslt := IsBalancedBrackets(s)

	if !rslt {
		t.Errorf("got %v, but expect %v", rslt, true)
	}
}
func TestIsBalancedBrackets_false(t *testing.T) {
	s := "{([)}"
	rslt := IsBalancedBrackets(s)

	if rslt {
		t.Errorf("got %v, but expect %v", rslt, false)
	}
}
