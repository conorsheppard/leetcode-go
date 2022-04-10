package add_two_numbers

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/add-two-numbers/

/*
You are given two non-empty linked lists, each representing a non-negative integer.
The numbers are stored in reverse order and each ListNode can contain only a single digit
(i.e. &ListNode{ Val:  2, Next: &ListNode{Val:  4, Next: &ListNode{Val:  9, Next: nil}}} == 942).
The answer should also be in reverse order i.e. [7,0,4,0,1] == 10,407.
Add the two numbers and return the sum as a linked list.
You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Constraints:
- The number of nodes in each linked list is in the range [1, 100].
- 0 <= Node.val <= 9
- It is guaranteed that the list represents a number that does not have leading zeros.
*/

func Test_AddTwoNumbers(t *testing.T) {
	tests := []struct {
		inputLinkedList1 *ListNode
		inputLinkedList2 *ListNode
		expectedResult   *ListNode
	}{
		// test 0: [7,0,8]
		{
			inputLinkedList1: &ListNode{
				Val:  2,
				Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: nil}},
			},
			inputLinkedList2: &ListNode{
				Val:  5,
				Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: nil}},
			},
			expectedResult: &ListNode{
				Val:  7,
				Next: &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: nil}},
			},
		},
		// test 1: [0]
		{
			inputLinkedList1: &ListNode{
				Val:  0,
				Next: nil,
			},
			inputLinkedList2: &ListNode{
				Val:  0,
				Next: nil,
			},
			expectedResult: &ListNode{
				Val:  0,
				Next: nil,
			},
		},
		// test 2: [8,9,9,9,0,0,0,1]
		{
			inputLinkedList1: &ListNode{
				Val:  9,
				Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}}}},
			},
			inputLinkedList2: &ListNode{
				Val:  9,
				Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}},
			},
			expectedResult: &ListNode{
				Val:  8,
				Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}}}},
			},
		},
		// test 3: [7,0,4,0,1]
		{
			inputLinkedList1: &ListNode{
				Val:  2,
				Next: &ListNode{Val: 4, Next: &ListNode{Val: 9, Next: nil}},
			},
			inputLinkedList2: &ListNode{
				Val:  5,
				Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9, Next: nil}}},
			},
			expectedResult: &ListNode{
				Val:  7,
				Next: &ListNode{Val: 0, Next: &ListNode{Val: 4, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}},
			},
		},
		// test 4: [6,6,4,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,1]
		{
			inputLinkedList1: &ListNode{
				Val:  1,
				Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}}}}}}}}}}}}}}}}}}}}}}}}}}},
			},
			inputLinkedList2: &ListNode{
				Val:  5,
				Next: &ListNode{Val: 6, Next: &ListNode{Val: 4, Next: &ListNode{Val: 9, Next: nil}}},
			},
			expectedResult: &ListNode{
				Val:  8,
				Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}}}},
			},
		},
		// test 5: [8,9,9,9,0,0,0,0,1]
		{
			inputLinkedList1: &ListNode{
				Val:  9,
				Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}}}}},
			},
			inputLinkedList2: &ListNode{
				Val:  9,
				Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: nil}}},
			},
			expectedResult: &ListNode{
				Val:  8,
				Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 9, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: nil}}}}}}},
			},
		},
	}

	for i := range tests {
		fmt.Printf("test %d\n", i)
		test := tests[i]

		result := AddTwoNumbers(test.inputLinkedList1, test.inputLinkedList2)

		temp := result

		resultStr := formatLinkedListAsString(temp)
		expectedResultStr := formatLinkedListAsString(test.expectedResult)

		fmt.Printf("result: %s\n", resultStr)
		fmt.Printf("expected result: %s\n", expectedResultStr)

		if expectedResultStr == resultStr {
			fmt.Println("PASS!!")
		} else {
			fmt.Println("FAIL!!")
		}
	}
}

func formatLinkedListAsString(list *ListNode) string {
	resultStr := ""
	for {
		if list.Next != nil {
			resultStr = fmt.Sprintf("%d%s", list.Val, resultStr)
			list = list.Next
		} else {
			resultStr = fmt.Sprintf("%d%s", list.Val, resultStr)
			break
		}
	}
	return resultStr
}
