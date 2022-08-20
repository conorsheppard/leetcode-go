package find_middle_node_linked_list

import l "github.com/conorsheppard/leetcode-go/internal/easy-problems/linked-list"

// https://leetcode.com/problems/middle-of-the-linked-list/
func middleNode(head *l.Node) *l.Node {
	middleNode, i := head, 0
	for ; head != nil; head, i = head.Next, i+1 {
		if i%2 == 1 {
			middleNode = middleNode.Next
		}
	}
	return middleNode
}

func middleNode2Pointers(head *l.Node) *l.Node {
	for fast := head; fast != nil && fast.Next != nil; head, fast = head.Next, fast.Next.Next {
	}
	return head
}
