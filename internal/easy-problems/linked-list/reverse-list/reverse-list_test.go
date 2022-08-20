package reverse_list

import (
	"fmt"
	l "github.com/conorsheppard/leetcode-go/internal/easy-problems/linked-list"
	"testing"
)

func Test_reverseList(t *testing.T) {
	tests := []struct {
		inputList            *l.Node
		expectedReversedList *l.Node
	}{
		{
			inputList:            &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: &l.Node{Val: 5, Next: &l.Node{Val: 6, Next: &l.Node{Val: 7, Next: &l.Node{Val: 8, Next: nil}}}}}}}},
			expectedReversedList: &l.Node{Val: 8, Next: &l.Node{Val: 7, Next: &l.Node{Val: 6, Next: &l.Node{Val: 5, Next: &l.Node{Val: 4, Next: &l.Node{Val: 3, Next: &l.Node{Val: 2, Next: &l.Node{Val: 1, Next: nil}}}}}}}},
		},
		{
			inputList:            nil,
			expectedReversedList: nil,
		},
	}

	for i, test := range tests {
		fmt.Printf("test: %d\n", i)

		result := reverseList(test.inputList)

		for ; test.expectedReversedList != nil; result, test.expectedReversedList = result.Next, test.expectedReversedList.Next {
			if result.Val != test.expectedReversedList.Val {
				t.Errorf("result.Val = %d, expectedReversedList.Val = %d", result.Val, test.expectedReversedList.Val)
			}
		}
	}
}
