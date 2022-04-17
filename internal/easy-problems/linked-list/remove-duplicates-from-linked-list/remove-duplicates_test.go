package remove_duplicates_from_linked_list

import (
	"fmt"
	l "leetcode/internal/easy-problems/linked-list"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		inputList    *l.ListNode
		expectedList *l.ListNode
	}{
		{
			inputList:    &l.ListNode{Val: 1, Next: &l.ListNode{Val: 1, Next: &l.ListNode{Val: 2, Next: &l.ListNode{Val: 2, Next: &l.ListNode{Val: 2, Next: nil}}}}},
			expectedList: &l.ListNode{Val: 1, Next: &l.ListNode{Val: 2, Next: nil}},
		},
		{
			inputList:    &l.ListNode{Val: 1, Next: &l.ListNode{Val: 1, Next: &l.ListNode{Val: 2, Next: &l.ListNode{Val: 2, Next: &l.ListNode{Val: 2, Next: &l.ListNode{Val: 3, Next: nil}}}}}},
			expectedList: &l.ListNode{Val: 1, Next: &l.ListNode{Val: 2, Next: &l.ListNode{Val: 3, Next: nil}}},
		},
		{
			inputList:    nil,
			expectedList: nil,
		},
	}

	for i, test := range tests {
		fmt.Printf("test: %d\n", i)

		result := deleteDuplicates(test.inputList)

		for ; test.expectedList != nil; result, test.expectedList = result.Next, test.expectedList.Next {
			if result.Val != test.expectedList.Val {
				t.Errorf("result.Val = %d, expectedList.Val = %d", result.Val, test.expectedList.Val)
				fmt.Printf("result list: %v\n", l.GetLinkedListAsString(result))
			}
		}
	}
}
