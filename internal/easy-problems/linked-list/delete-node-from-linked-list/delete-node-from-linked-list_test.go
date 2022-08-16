package delete_node_from_linked_list

import (
	"fmt"
	l "leetcode-go/internal/easy-problems/linked-list"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	tests := []struct {
		linkedList     *l.Node
		deleteNodeVal  int
		expectedResult *l.Node
	}{
		{
			linkedList:     &l.Node{Val: 4, Next: &l.Node{Val: 5, Next: &l.Node{Val: 1, Next: &l.Node{Val: 9, Next: nil}}}},
			deleteNodeVal:  5,
			expectedResult: &l.Node{Val: 4, Next: &l.Node{Val: 1, Next: &l.Node{Val: 9, Next: nil}}},
		},
		{
			linkedList:     &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: nil}},
			deleteNodeVal:  1,
			expectedResult: &l.Node{Val: 2, Next: nil},
		},
		{
			linkedList:     &l.Node{Val: 3, Next: &l.Node{Val: 5, Next: &l.Node{Val: 1, Next: &l.Node{Val: 7, Next: &l.Node{Val: 8, Next: nil}}}}},
			deleteNodeVal:  1,
			expectedResult: &l.Node{Val: 3, Next: &l.Node{Val: 5, Next: &l.Node{Val: 7, Next: &l.Node{Val: 8, Next: nil}}}},
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d\n", i)

		var ref *l.Node
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
