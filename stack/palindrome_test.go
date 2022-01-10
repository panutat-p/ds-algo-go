package stack

import "testing"

func TestIsPalindrome(t *testing.T) {
	str := "abc|cba"
	rslt := IsPalindrome(str)

	if !rslt {
		t.Errorf("got %v, but expect %v", rslt, true)
	}
}

func TestIsPalindrome2(t *testing.T) {
	str := "I did, did I?"
	rslt := IsPalindrome(str)

	if !rslt {
		t.Errorf("got %v, but expect %v", rslt, true)
	}
}

func TestIsPalindrome3(t *testing.T) {
	str := "Don't nod"
	rslt := IsPalindrome(str)

	if !rslt {
		t.Errorf("got %v, but expect %v", rslt, true)
	}
}

func TestIsPalindrome4(t *testing.T) {
	str := "hello"
	rslt := IsPalindrome(str)

	if rslt {
		t.Errorf("got %v, but expect %v", rslt, false)
	}
}
