package remove_duplicates_from_linked_list

import l "leetcode/internal/easy-problems/linked-list"

func deleteDuplicates(head *l.Node) *l.Node {
	if head == nil {
		return head
	}
	ref := head
	for head.Next != nil {
		for head.Next != nil && head.Val == head.Next.Val {
			*head = *head.Next
		}
		if head.Next != nil {
			head = head.Next
		}
	}
	return ref
}

func deleteDuplicatesIfElse(head *l.Node) *l.Node {
	if head == nil {
		return head
	}
	ref := head
	for head.Next != nil {
		if head.Val == head.Next.Val {
			*head = *head.Next
		} else {
			head = head.Next
		}
	}
	return ref
}
