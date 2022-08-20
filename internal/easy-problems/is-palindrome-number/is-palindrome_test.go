package is_palindrome_number

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/palindrome-number/
func Test_IsPalindrome(t *testing.T) {
	tests := []struct {
		input        int
		isPalindrome bool
	}{
		{0, true},
		{11, true},
		{121, true},
		{-121, false},
		{421245, false},
		{33333, true},
		{10, false},
		{7777778, false},
		{1234567654321, true},
		{444445444444, false},
	}

	for i := range tests {
		test := tests[i]
		fmt.Printf("test %d\n", i)

		result := isPalindrome(test.input)

		if result != test.isPalindrome {
			if result {
				t.Errorf("%d is not a palindrome but result was %t", test.input, result)
			} else {
				t.Errorf("%d is a palindrome but result was %t", test.input, result)
			}
		}
	}
}
