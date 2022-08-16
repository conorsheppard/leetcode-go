package remove_linked_list_elements_203

import l "leetcode-go/internal/easy-problems/linked-list"

// https://leetcode.com/problems/remove-linked-list-elements/
func removeElements(head *l.Node, val int) *l.Node {
	if head == nil {
		return nil
	}
	head = &l.Node{Val: -1, Next: head}
	ref, slow := head.Next, head
	for ref.Next != nil {
		if ref.Val == val {
			*ref = *ref.Next
		} else {
			ref, slow = ref.Next, slow.Next
		}
	}

	if ref.Val == val {
		slow.Next = nil
	}

	return head.Next
}
