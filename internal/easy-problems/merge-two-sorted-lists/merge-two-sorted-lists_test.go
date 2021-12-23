package merge_two_sorted_lists

import (
	"fmt"
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
	tests := []struct{
		list1 *ListNode
		list2 *ListNode
		expectedList *ListNode
	}{
		{
			list1: nil,
			list2: nil,
			expectedList: nil,
		},
		{
			list1: nil,
			list2: &ListNode{Val:  0, Next: nil},
			expectedList: &ListNode{Val:  0, Next: nil},
		},
		{
			list1: &ListNode{
				Val:  2,
				Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: nil}},
			},
			list2: &ListNode{
				Val:  5,
				Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}},
			},
			expectedList: &ListNode{
				Val:  2,
				Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val:  5,
					Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}},
				}}},
			},
		},
		{
			list1: &ListNode{
				Val:  2,
				Next: &ListNode{Val: 3, Next: &ListNode{Val: 7, Next: nil}},
			},
			list2: &ListNode{
				Val:  5,
				Next: &ListNode{Val: 6, Next: &ListNode{Val: 8, Next: nil}},
			},
			expectedList: &ListNode{
				Val:  2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next:
				&ListNode{Val:  6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 8, Next: nil}},
				}}},
			},
		},
		{
			list1: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7, Next: nil}}},
			},
			list2: &ListNode{
				Val:  5,
				Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}},
			},
			expectedList: &ListNode{
				Val:  1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next:
				&ListNode{Val:  5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: nil}}},
				}}},
			},
		},
		{
			list1: &ListNode{
				Val:  7,
				Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: nil}}},
			},
			list2: &ListNode{
				Val:  7,
				Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: nil}},
			},
			expectedList: &ListNode{
				Val:  7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next:
				&ListNode{Val:  7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: &ListNode{Val: 7, Next: nil}}},
				}}},
			},
		},
		{
			list1: &ListNode{Val:  1, Next: nil},
			list2: &ListNode{Val:  0, Next: nil},
			expectedList: &ListNode{Val:  0, Next: &ListNode{Val:  1, Next: nil}},
		},
	}

	for i := range tests {
		test := tests[i]
		fmt.Printf("test %d\n", i)

		result := mergeTwoLists(test.list1, test.list2)

		for result != nil || test.expectedList != nil {
			if result == nil && test.expectedList != nil {
				t.Error("not equal, result is nil, test.expectedList not nil")
				break
			} else if result != nil && test.expectedList == nil {
				t.Errorf("not equal, result is not nil, test.expectedList is nil")
				t.Skip()
			} else if result == nil && test.expectedList == nil {
				fmt.Println("lists are both nil. PASS.")
				break
			}
			if result.Val != test.expectedList.Val {
				t.Error("values are not equal")
				break
			}
			result = result.Next
			test.expectedList = test.expectedList.Next
		}
	}
}