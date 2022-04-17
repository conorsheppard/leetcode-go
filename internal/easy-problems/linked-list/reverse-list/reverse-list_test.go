package reverse_list

import (
	"fmt"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		inputList            *ListNode
		expectedReversedList *ListNode
	}{
		{
			inputList:            &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: nil}}}}}}}},
			expectedReversedList: &ListNode{Val: 8, Next: &ListNode{Val: 7, Next: &ListNode{Val: 6, Next: &ListNode{Val: 5, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}}}}},
		},
		{
			inputList:            nil,
			expectedReversedList: nil,
		},
	}

	for i, test := range tests {
		fmt.Printf("test: %d\n", i)

		result := reverseList(test.inputList)

		for ; test.expectedReversedList != nil; result, test.expectedReversedList = result.Next, test.expectedReversedList.Next {
			if result.Val != test.expectedReversedList.Val {
				t.Errorf("result.Val = %d, expectedReversedList.Val = %d", result.Val, test.expectedReversedList.Val)
			}
		}
	}
}
