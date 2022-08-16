package linked_list_cycle_I

import l "leetcode-go/internal/easy-problems/linked-list"

func hasCycle(head *l.Node) bool {
	var addressMap = make(map[*l.Node]*int)
	for ; head != nil; head = head.Next {
		if addressMap[head] == &head.Val {
			return true
		}
		addressMap[head] = &head.Val
	}
	return false
}
