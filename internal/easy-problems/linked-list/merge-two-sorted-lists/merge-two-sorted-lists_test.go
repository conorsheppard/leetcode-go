package merge_two_sorted_lists

import (
	"fmt"
	l "github.com/conorsheppard/leetcode-go/internal/easy-problems/linked-list"
	"testing"
)

// You are given the heads of two sorted linked lists list1 and list2.
// Merge the two lists into one sorted list. The list should be made by splicing together the nodes of the first two lists.
// Return the head of the merged linked list.

// Constraints:
// The number of nodes in both lists is in the range [0, 50].
// -100 <= Node.val <= 100
// Both list1 and list2 are sorted in non-decreasing order.

func Test_MergeTwoSortedLists(t *testing.T) {
	tests := []struct {
		list1        *l.Node
		list2        *l.Node
		expectedList *l.Node
	}{
		{
			list1:        nil,
			list2:        nil,
			expectedList: nil,
		},
		{
			list1:        nil,
			list2:        &l.Node{Val: 0, Next: nil},
			expectedList: &l.Node{Val: 0, Next: nil},
		},
		{
			list1: &l.Node{
				Val:  2,
				Next: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: nil}},
			},
			list2: &l.Node{
				Val:  5,
				Next: &l.Node{Val: 6, Next: &l.Node{Val: 7, Next: nil}},
			},
			expectedList: &l.Node{
				Val: 2,
				Next: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: &l.Node{Val: 5,
					Next: &l.Node{Val: 6, Next: &l.Node{Val: 7, Next: nil}},
				}}},
			},
		},
		{
			list1: &l.Node{
				Val:  2,
				Next: &l.Node{Val: 3, Next: &l.Node{Val: 7, Next: nil}},
			},
			list2: &l.Node{
				Val:  5,
				Next: &l.Node{Val: 6, Next: &l.Node{Val: 8, Next: nil}},
			},
			expectedList: &l.Node{
				Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 5, Next: &l.Node{Val: 6, Next: &l.Node{Val: 7, Next: &l.Node{Val: 8, Next: nil}}}}},
			},
		},
		{
			list1: &l.Node{
				Val:  1,
				Next: &l.Node{Val: 2, Next: &l.Node{Val: 5, Next: &l.Node{Val: 7, Next: nil}}},
			},
			list2: &l.Node{
				Val:  5,
				Next: &l.Node{Val: 6, Next: &l.Node{Val: 7, Next: nil}},
			},
			expectedList: &l.Node{
				Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 5, Next: &l.Node{Val: 5, Next: &l.Node{Val: 6, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: nil}}}}}},
			},
		},
		{
			list1: &l.Node{
				Val:  7,
				Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: nil}}},
			},
			list2: &l.Node{
				Val:  7,
				Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: nil}},
			},
			expectedList: &l.Node{
				Val: 7, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: nil}}}}}},
			},
		},
		{
			list1:        &l.Node{Val: 1, Next: nil},
			list2:        &l.Node{Val: 0, Next: nil},
			expectedList: &l.Node{Val: 0, Next: &l.Node{Val: 1, Next: nil}},
		},
	}

	for i := range tests {
		test := tests[i]
		fmt.Printf("test %d\n", i)

		result := mergeTwoLists(test.list1, test.list2)

		success, message := l.CheckListEquality(result, test.expectedList)

		if !success {
			t.Errorf("FAIL: %s", message)
		} else {
			fmt.Printf(message)
		}
	}
}
