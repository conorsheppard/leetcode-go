package delete_node_from_linked_list

import (
	"fmt"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	tests := []struct {
		linkedList     *ListNode
		deleteNodeVal  int
		expectedResult *ListNode
	}{
		{
			linkedList:     &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: nil}}}},
			deleteNodeVal:  5,
			expectedResult: &ListNode{Val: 4, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9, Next: nil}}},
		},
		{
			linkedList:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			deleteNodeVal:  1,
			expectedResult: &ListNode{Val: 2, Next: nil},
		},
		{
			linkedList:     &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: nil}}}}},
			deleteNodeVal:  1,
			expectedResult: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: nil}}}},
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d\n", i)

		var ref *ListNode
		linkedList := test.linkedList

		for ; linkedList.Next != nil; linkedList = linkedList.Next {
			if linkedList.Val == test.deleteNodeVal {
				ref = linkedList
			}
		}

		deleteNode(ref)

		for ; test.linkedList != nil; test.linkedList, test.expectedResult = test.linkedList.Next, test.expectedResult.Next {
			if test.linkedList.Val != test.expectedResult.Val {
				t.Errorf("test.inputList.Val: %d != test.expectedResult.Val: %d", test.linkedList.Val, test.expectedResult.Val)

			}
		}
	}
}
