package list_intersection_160

import (
	l "leetcode/internal/easy-problems/linked-list"
	"math"
)

// https://leetcode.com/problems/intersection-of-two-linked-lists/
func getIntersectionNode(headA, headB *l.Node) *l.Node {
	lenHeadA, lenHeadB, newHeadA, newHeadB := 1, 1, headA, headB
	for ref1 := headA; ref1 != nil && ref1.Next != nil; ref1 = ref1.Next {
		lenHeadA++
	}

	for ref2 := headB; ref2 != nil && ref2.Next != nil; ref2 = ref2.Next {
		lenHeadB++
	}
	offset := int(math.Abs(float64(lenHeadA) - float64(lenHeadB)))
	if lenHeadA > lenHeadB {
		for j := 0; j < offset; j, newHeadA = j+1, newHeadA.Next {
		}
	} else {
		for j := 0; j < offset; j, newHeadB = j+1, newHeadB.Next {
		}
	}

	for ; newHeadA.Next != newHeadB.Next; newHeadA, newHeadB = newHeadA.Next, newHeadB.Next {
	}
	if newHeadA == newHeadB {
		return newHeadA
	}

	return newHeadA.Next
}
