package is_palindrome_string

import (
	"fmt"
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	tests := []struct {
		input          string
		expectedResult bool
	}{
		{"hello", false},
		{"a", true},
		{"hannah", true},
		{"", true},
		{"Anna", true},
	}

	for i, test := range tests {
		fmt.Printf("test %d\n", i)
		result := isPalindrome(test.input)
		if result != test.expectedResult {
			t.Errorf("FAIL: expected %t, got %t", test.expectedResult, result)
		} else {
			fmt.Printf("\tPASS\n")
		}
	}
}
