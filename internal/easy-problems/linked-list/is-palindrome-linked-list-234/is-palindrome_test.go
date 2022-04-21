package is_palindrome_linked_list_234

import (
	"fmt"
	l "leetcode/internal/easy-problems/linked-list"
	"testing"
)

func Test_IsPalindrome(t *testing.T) {
	tests := []struct {
		input          *l.Node
		expectedResult bool
	}{
		{
			input: &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: &l.Node{Val: 4, Next: &l.Node{Val: 4, Next: &l.Node{Val: 5,
				Next: &l.Node{Val: 4, Next: &l.Node{Val: 4, Next: &l.Node{Val: 4, Next: &l.Node{Val: 3, Next: &l.Node{Val: 2, Next: &l.Node{Val: 1, Next: nil}}}}}}}}}}}}},
			expectedResult: true,
		},
		{
			input:          &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 1, Next: nil}}},
			expectedResult: true,
		},
		{
			input:          &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 2, Next: &l.Node{Val: 1, Next: nil}}}},
			expectedResult: true,
		},
		{
			input:          &l.Node{Val: 2, Next: nil},
			expectedResult: true,
		},
		{
			input:          &l.Node{Val: 2, Next: &l.Node{Val: 2, Next: nil}},
			expectedResult: true,
		},
		{
			input:          &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 2, Next: &l.Node{Val: 1, Next: nil}}}}},
			expectedResult: true,
		},
		{
			input:          &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: nil}}},
			expectedResult: false,
		},
		{
			input:          &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 2, Next: &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 2, Next: &l.Node{Val: 1, Next: nil}}}}}}},
			expectedResult: true,
		},
		{
			input:          &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 3, Next: &l.Node{Val: 5, Next: &l.Node{Val: 3, Next: &l.Node{Val: 3, Next: &l.Node{Val: 2, Next: &l.Node{Val: 1, Next: &l.Node{Val: 1, Next: nil}}}}}}}}}},
			expectedResult: false,
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d\n", i)
		result := isPalindrome(test.input)

		if result != test.expectedResult {
			t.Errorf("FAIL: expected: %t, got: %t", test.expectedResult, result)
		}
	}
}
