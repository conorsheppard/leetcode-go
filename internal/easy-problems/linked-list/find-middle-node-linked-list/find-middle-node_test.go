package find_middle_node_linked_list

import (
	"fmt"
	l "github.com/conorsheppard/leetcode-go/internal/easy-problems/linked-list"
	"testing"
)

// Given the head of a singly linked list, return the middle node of the linked list.
// If there are two middle nodes, return the second middle node.
// https://leetcode.com/problems/middle-of-the-linked-list/
func Test_Convert(t *testing.T) {
	tests := []struct {
		inputList          *l.Node
		expectedMiddleNode *l.Node
	}{
		//{ removed because problem description says list is not empty
		//	inputList:      nil,
		//	expectedResult: 0,
		//},
		{
			inputList:          &l.Node{Val: 1, Next: nil},
			expectedMiddleNode: &l.Node{Val: 1, Next: nil},
		},
		{
			inputList:          &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: nil}},
			expectedMiddleNode: &l.Node{Val: 2, Next: nil},
		},
		{
			inputList:          &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: nil}}},
			expectedMiddleNode: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: nil}},
		},
		{
			inputList:          &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: nil}}}},
			expectedMiddleNode: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: nil}},
		},
		{
			inputList:          &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: &l.Node{Val: 5, Next: &l.Node{Val: 6, Next: &l.Node{Val: 7, Next: &l.Node{Val: 8, Next: nil}}}}}}}},
			expectedMiddleNode: &l.Node{Val: 5, Next: &l.Node{Val: 6, Next: &l.Node{Val: 7, Next: &l.Node{Val: 8, Next: nil}}}},
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
