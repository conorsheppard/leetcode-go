package linked_list_cycle_I

import l "leetcode/internal/easy-problems/linked-list"

func hasCycle(head *l.ListNode) *l.ListNode {
	var addressMap = make(map[*l.ListNode]*int)

	for ; head != nil; head = head.Next {
		if addressMap[head] != nil {
			return head
		}
		addressMap[head] = &head.Val
	}
	return nil
}
