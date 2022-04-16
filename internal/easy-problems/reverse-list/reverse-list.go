package reverse_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var cur *ListNode
	for ; head != nil; head = head.Next {
		newNode := &ListNode{Val: head.Val, Next: cur}
		cur = newNode
	}
	return cur
}
