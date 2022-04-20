package reverse_list

import l "leetcode/internal/easy-problems/linked-list"

func reverseList(head *l.Node) *l.Node {
	var cur *l.Node
	for ; head != nil; head = head.Next {
		newNode := &l.Node{Val: head.Val, Next: cur}
		cur = newNode
	}
	return cur
}
