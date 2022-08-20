package linked_list_cycle_I

import (
	"fmt"
	l "github.com/conorsheppard/leetcode-go/internal/easy-problems/linked-list"
	"testing"
)

func Test_HasCycle(t *testing.T) {
	tests := []struct {
		inputList      *l.Node
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
			inputList:      &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: nil}}},
			expectedResult: false,
			cycleFrom:      -1,
		},
		{
			inputList:      &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 5, Next: &l.Node{Val: 7, Next: &l.Node{Val: 8, Next: nil}}}}},
			expectedResult: true,
			cycleFrom:      4,
			cycleTo:        2,
		},
		{
			inputList:      &l.Node{Val: 1, Next: nil},
			expectedResult: false,
			cycleFrom:      -1,
		},
		{
			inputList:      &l.Node{Val: 1, Next: nil},
			expectedResult: true,
			cycleFrom:      0,
			cycleTo:        0,
		},
		{
			inputList:      &l.Node{Val: 0, Next: &l.Node{Val: 0, Next: nil}},
			expectedResult: false,
			cycleFrom:      -1,
		},
		{
			inputList:      &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: nil}},
			expectedResult: true,
			cycleFrom:      1,
			cycleTo:        0,
		},
		{
			inputList:      &l.Node{Val: 3, Next: &l.Node{Val: 2, Next: &l.Node{Val: 0, Next: &l.Node{Val: -4, Next: nil}}}},
			expectedResult: true,
			cycleFrom:      3,
			cycleTo:        1,
		},
		{
			inputList:      &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 5, Next: &l.Node{Val: 7, Next: &l.Node{Val: 8, Next: &l.Node{Val: 0, Next: nil}}}}}},
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

func addCycle(head *l.Node, cycleFrom, cycleTo int) {
	if cycleFrom == -1 {
		return
	}
	var targetNode *l.Node
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
