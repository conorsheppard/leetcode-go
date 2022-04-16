package find_middle_node_linked_list

import (
	"fmt"
	"testing"
)

// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.
// https://leetcode.com/problems/middle-of-the-linked-list/
func Test_Convert(t *testing.T) {
	tests := []struct {
		inputList          *ListNode
		expectedMiddleNode *ListNode
	}{
		//{ removed because problem description says list is not empty
		//	inputList:      nil,
		//	expectedResult: 0,
		//},
		{
			inputList:          &ListNode{Val: 1, Next: nil},
			expectedMiddleNode: &ListNode{Val: 1, Next: nil},
		},
		{
			inputList:          &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			expectedMiddleNode: &ListNode{Val: 2, Next: nil},
		},
		{
			inputList:          &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}},
			expectedMiddleNode: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}},
		},
		{
			inputList:          &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}}}},
			expectedMiddleNode: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}},
		},
		{
			inputList:          &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: nil}}}}}}}},
			expectedMiddleNode: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: nil}}}},
		},
	}

	for i, test := range tests {
		fmt.Printf("test: %d\n", i)

		result := middleNode(test.inputList)

		if result.Val != test.expectedMiddleNode.Val {
			t.Errorf("result = %d, expecetdResult = %d", result.Val, test.expectedMiddleNode.Val)
		}
	}
}
