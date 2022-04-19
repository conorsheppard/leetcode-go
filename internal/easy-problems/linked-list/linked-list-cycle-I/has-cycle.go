package linked_list_cycle_I

import l "leetcode/internal/easy-problems/linked-list"

func hasCycle(head *l.ListNode) bool {
	var addressMap = make(map[*l.ListNode]*int)
	for ; head != nil; head = head.Next {
		if addressMap[head] == &head.Val {
			return true
		}
		addressMap[head] = &head.Val
	}
	return false
}
