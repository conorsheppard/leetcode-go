package list_intersection_160

import (
	"fmt"
	l "leetcode/internal/easy-problems/linked-list"
	"testing"
)

var listTest2 = &l.Node{Val: 1, Next: nil}

// https://leetcode.com/problems/intersection-of-two-linked-lists/
func Test_GetIntersectionNode(t *testing.T) {
	tests := []struct {
		headA        *l.Node
		headB        *l.Node
		intersectVal int // The value of the node where the intersection occurs. This is 0 if there is no intersected node.
	}{
		{
			headA:        &l.Node{Val: 3, Next: nil},
			headB:        &l.Node{Val: 2, Next: &l.Node{Val: 3, Next: nil}},
			intersectVal: 3,
		},
		{
			headA:        listTest2,
			headB:        listTest2,
			intersectVal: 1,
		},
		{
			headA:        &l.Node{Val: 1, Next: &l.Node{Val: 9, Next: &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 4, Next: nil}}}}},
			headB:        &l.Node{Val: 3, Next: nil},
			intersectVal: 2,
		},
		{
			headA:        &l.Node{Val: 4, Next: &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: &l.Node{Val: 8, Next: &l.Node{Val: 4, Next: &l.Node{Val: 5, Next: nil}}}}}},
			headB:        &l.Node{Val: 9, Next: &l.Node{Val: 7, Next: &l.Node{Val: 2, Next: &l.Node{Val: 5, Next: &l.Node{Val: 6, Next: &l.Node{Val: 1, Next: &l.Node{Val: 2, Next: nil}}}}}}},
			intersectVal: 8,
		},
		{
			headA:        &l.Node{Val: 2, Next: &l.Node{Val: 6, Next: &l.Node{Val: 4, Next: nil}}},
			headB:        &l.Node{Val: 1, Next: &l.Node{Val: 5, Next: nil}},
			intersectVal: 0,
		},
		{
			headA:        &l.Node{Val: 0, Next: nil},
			headB:        &l.Node{Val: 0, Next: nil},
			intersectVal: 0,
		},
		{
			headA:        &l.Node{Val: 4, Next: &l.Node{Val: 1, Next: &l.Node{Val: 8, Next: &l.Node{Val: 4, Next: &l.Node{Val: 5, Next: nil}}}}},
			headB:        &l.Node{Val: 5, Next: &l.Node{Val: 6, Next: &l.Node{Val: 1, Next: nil}}},
			intersectVal: 0,
		},
		{
			headA:        &l.Node{Val: 4, Next: &l.Node{Val: 1, Next: &l.Node{Val: 8, Next: &l.Node{Val: 4, Next: &l.Node{Val: 5, Next: nil}}}}},
			headB:        &l.Node{Val: 5, Next: &l.Node{Val: 6, Next: &l.Node{Val: 1, Next: nil}}},
			intersectVal: 8,
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d\n", i)

		interSectionNode := addIntersection(test.headA, test.headB, test.intersectVal)

		result := getIntersectionNode(test.headA, test.headB)

		if result != interSectionNode {
			t.Errorf("FAIL: result.Val: %v, test.intersectVal: %v", result, interSectionNode)
		} else {
			fmt.Printf("PASS!!\n")
		}
	}
}

func addIntersection(headA, headB *l.Node, intersectVal int) *l.Node {
	if headA == headB {
		return headA
	}
	if intersectVal == 0 {
		return nil
	}
	refA, refB := headA, headB
	for ; refB != nil && refB.Next != nil; refB = refB.Next {
	}

	for ; intersectVal != refA.Val; refA = refA.Next {
	}
	refB.Next = refA

	return refA
}
