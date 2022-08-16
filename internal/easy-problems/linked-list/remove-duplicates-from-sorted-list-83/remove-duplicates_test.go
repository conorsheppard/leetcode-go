package remove_duplicates_from_sorted_list_83

import (
	"fmt"
	l "leetcode-go/internal/easy-problems/linked-list"
	"testing"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-list/
func Test_reverseList(t *testing.T) {
	tests := []struct {
		inputList    *l.Node
		expectedList *l.Node
	}{
		{
			inputList:    &l.Node{Val: 1, Next: &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 2, Next: &l.Node{Val: 2, Next: nil}}}}},
			expectedList: &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: nil}},
		},
		{
			inputList:    &l.Node{Val: 1, Next: &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 2, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: nil}}}}}},
			expectedList: &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: nil}}},
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
