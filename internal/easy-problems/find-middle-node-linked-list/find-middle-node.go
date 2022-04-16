package find_middle_node_linked_list

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// https://leetcode.com/problems/middle-of-the-linked-list/
func middleNode(head *ListNode) *ListNode {
	middleNode, i := head, 0
	for ; head != nil; head, i = head.Next, i+1 {
		if i%2 == 1 {
			middleNode = middleNode.Next
		}
	}
	return middleNode
}
