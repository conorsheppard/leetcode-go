package convert_binary_linked_list

import (
	"fmt"
	"testing"
)

// The most significant bit is the first node in the list i.e. [1,0,1,1] == 11
// Return the decimal value of the number in a linked list-represented binary number.
func Test_Convert(t *testing.T) {
	tests := []struct {
		inputList      *ListNode
		expectedResult int
	}{
		//{ removed because problem description says list is not empty
		//	inputList:      nil,
		//	expectedResult: 0,
		//},
		{
			inputList:      &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}},
			expectedResult: 3,
		},
		{
			inputList:      &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}},
			expectedResult: 5,
		},
		{
			inputList:      &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}}}},
			expectedResult: 11,
		},
		{
			inputList:      &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}}}}}}}},
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
