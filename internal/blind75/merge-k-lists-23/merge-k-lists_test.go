package merge_k_lists_23

import (
	"fmt"
	l "leetcode/internal/easy-problems/linked-list"
	"testing"
)

func Test_MergeKLists(t *testing.T) {
	tests := []struct {
		inputLists     []*l.Node
		expectedResult *l.Node
	}{
		{
			inputLists:     []*l.Node{},
			expectedResult: nil,
		},
		{
			inputLists: []*l.Node{
				{Val: 1, Next: &l.Node{Val: 4, Next: &l.Node{Val: 5, Next: nil}}},
				{Val: 1, Next: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: nil}}},
				{Val: 2, Next: &l.Node{Val: 6, Next: nil}},
			},
			expectedResult: &l.Node{Val: 1, Next: &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: &l.Node{Val: 4, Next: &l.Node{Val: 4, Next: &l.Node{Val: 5, Next: &l.Node{Val: 6, Next: nil}}}}}}}},
		},
	}

	var pass, fail int

	for i, test := range tests {
		fmt.Printf("test %d\n", i)

		result := mergeKLists(test.inputLists)

		if fmt.Sprintf("%v\n", result) != fmt.Sprintf("%v\n", test.expectedResult) {
			t.Errorf("FAIL: test %d\n", i)
			t.Errorf("Expected: %s\nGot: %s\n", l.GetLinkedListAsString(test.expectedResult), l.GetLinkedListAsString(result))
			fail++
		} else {
			pass++
		}
	}
	fmt.Printf("Tests ran: %d\n", len(tests))
	fmt.Printf("Tests passed: %d\n", pass)
	fmt.Printf("Tests failed: %d\n", fail)
}
