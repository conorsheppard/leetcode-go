package convert_binary_linked_list

import (
	"fmt"
	l "github.com/conorsheppard/leetcode-go/internal/easy-problems/linked-list"
	"testing"
)

// The most significant bit is the first node in the list i.e. [1,0,1,1] == 11
// Return the decimal value of the number in a linked list-represented binary number.
func Test_Convert(t *testing.T) {
	tests := []struct {
		inputList      *l.Node
		expectedResult int
	}{
		//{ removed because problem description says list is not empty
		//	inputList:      nil,
		//	expectedResult: 0,
		//},
		{
			inputList:      &l.Node{Val: 1, Next: &l.Node{Val: 1, Next: nil}},
			expectedResult: 3,
		},
		{
			inputList:      &l.Node{Val: 1, Next: &l.Node{Val: 0, Next: &l.Node{Val: 1, Next: nil}}},
			expectedResult: 5,
		},
		{
			inputList:      &l.Node{Val: 1, Next: &l.Node{Val: 0, Next: &l.Node{Val: 1, Next: &l.Node{Val: 1, Next: nil}}}},
			expectedResult: 11,
		},
		{
			inputList:      &l.Node{Val: 1, Next: &l.Node{Val: 0, Next: &l.Node{Val: 1, Next: &l.Node{Val: 1, Next: &l.Node{Val: 1, Next: &l.Node{Val: 0, Next: &l.Node{Val: 1, Next: &l.Node{Val: 1, Next: nil}}}}}}}},
			expectedResult: 187,
		},
	}

	for i, test := range tests {
		fmt.Printf("test: %d\n", i)

		result := getDecimalValue(test.inputList)

		if result != test.expectedResult {
			t.Errorf("result = %d, expecetdResult = %d", result, test.expectedResult)
		}
	}
}
