package remove_linked_list_elements_203

import (
	"fmt"
	l "leetcode/internal/easy-problems/linked-list"
	"testing"
)

func Test_RemoveElements(t *testing.T) {
	tests := []struct {
		input          *l.Node
		removeVal      int
		expectedOutput *l.Node
	}{
		{
			input:          &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: &l.Node{Val: 1, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: nil}}}}}}},
			removeVal:      7,
			expectedOutput: &l.Node{Val: 1, Next: nil},
		},
		{
			input:          &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: &l.Node{Val: 7, Next: nil}}}},
			removeVal:      7,
			expectedOutput: nil,
		},
		{
			input:          &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 6, Next: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: &l.Node{Val: 5, Next: &l.Node{Val: 6, Next: nil}}}}}}},
			removeVal:      6,
			expectedOutput: &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: &l.Node{Val: 5, Next: nil}}}}},
		},
		{
			input:          nil,
			removeVal:      1,
			expectedOutput: nil,
		},
		{
			input:          &l.Node{Val: 7, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 7, Next: nil}}}},
			removeVal:      0,
			expectedOutput: &l.Node{Val: 7, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 7, Next: nil}}}},
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d\n", i)

		result := removeElements(test.input, test.removeVal)

		success, message := checkListEquality(result, test.expectedOutput)

		if !success {
			t.Errorf("FAIL: %s", message)
		} else {
			fmt.Printf(message)
		}
	}
}

func checkListEquality(result, expectedList *l.Node) (bool, string) {
	for result != nil || expectedList != nil {
		if result == nil && expectedList != nil {
			return false, "not equal, result is nil, test.expectedList not nil\n"
		} else if result != nil && expectedList == nil {
			return false, "not equal, result is not nil, test.expectedList is nil\n"
		} else if result == nil && expectedList == nil {
			return true, "lists are both nil. PASS.\n"
		}
		if result.Val != expectedList.Val {
			return false, "expected: " + l.GetLinkedListAsString(expectedList) + ", got: " + l.GetLinkedListAsString(result) + "\n"
		}
		result = result.Next
		expectedList = expectedList.Next
	}

	return true, "PASS\n"
}
