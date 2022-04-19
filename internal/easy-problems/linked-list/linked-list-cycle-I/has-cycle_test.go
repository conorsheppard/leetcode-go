package linked_list_cycle_I

import (
	"fmt"
	l "leetcode/internal/easy-problems/linked-list"
	"testing"
)

func Test_HasCycle(t *testing.T) {
	tests := []struct {
		inputList      *l.ListNode
		expectedResult bool
		cycleFrom      int
		cycleTo        int
	}{
		{
			inputList:      nil,
			expectedResult: false,
			cycleFrom:      -1,
		},
		{
			inputList:      &l.ListNode{Val: 2, Next: &l.ListNode{Val: 3, Next: &l.ListNode{Val: 4, Next: nil}}},
			expectedResult: false,
			cycleFrom:      -1,
		},
		{
			inputList:      &l.ListNode{Val: 1, Next: &l.ListNode{Val: 2, Next: &l.ListNode{Val: 5, Next: &l.ListNode{Val: 7, Next: &l.ListNode{Val: 8, Next: nil}}}}},
			expectedResult: true,
			cycleFrom:      4,
			cycleTo:        2,
		},
		{
			inputList:      &l.ListNode{Val: 1, Next: nil},
			expectedResult: false,
			cycleFrom:      -1,
		},
		{
			inputList:      &l.ListNode{Val: 1, Next: nil},
			expectedResult: true,
			cycleFrom:      0,
			cycleTo:        0,
		},
		{
			inputList:      &l.ListNode{Val: 0, Next: &l.ListNode{Val: 0, Next: nil}},
			expectedResult: false,
			cycleFrom:      -1,
		},
		{
			inputList:      &l.ListNode{Val: 1, Next: &l.ListNode{Val: 2, Next: nil}},
			expectedResult: true,
			cycleFrom:      1,
			cycleTo:        0,
		},
		{
			inputList:      &l.ListNode{Val: 3, Next: &l.ListNode{Val: 2, Next: &l.ListNode{Val: 0, Next: &l.ListNode{Val: -4, Next: nil}}}},
			expectedResult: true,
			cycleFrom:      3,
			cycleTo:        1,
		},
		{
			inputList:      &l.ListNode{Val: 1, Next: &l.ListNode{Val: 2, Next: &l.ListNode{Val: 5, Next: &l.ListNode{Val: 7, Next: &l.ListNode{Val: 8, Next: &l.ListNode{Val: 0, Next: nil}}}}}},
			expectedResult: false,
			cycleFrom:      -1,
		},
	}

	for i := range tests {
		test := tests[i]
		fmt.Printf("test %d\n", i)

		addCycle(test.inputList, test.cycleFrom, test.cycleTo)

		result := hasCycle(test.inputList)

		if result != test.expectedResult {
			t.Errorf("result = %t, expecetdResult = %t", result, test.expectedResult)
		}
	}
}

func addCycle(head *l.ListNode, cycleFrom, cycleTo int) {
	if cycleFrom == -1 {
		return
	}
	var targetNode *l.ListNode
	for i := 0; i <= cycleFrom; i++ {
		if i == cycleTo {
			targetNode = head
		}
		if i == cycleFrom {
			head.Next = targetNode
		}
		head = head.Next
	}
}
