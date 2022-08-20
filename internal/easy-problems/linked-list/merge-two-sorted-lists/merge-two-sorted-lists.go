package merge_two_sorted_lists

import l "github.com/conorsheppard/leetcode-go/internal/easy-problems/linked-list"

// https://leetcode.com/problems/merge-two-sorted-lists/
func mergeTwoLists(list1, list2 *l.Node) *l.Node {
	curr := &l.Node{}
	res := curr
	for list1 != nil || list2 != nil {
		if list1 == nil {
			curr.Next = list2
			break
		}
		if list2 == nil {
			curr.Next = list1
			break
		}

		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	return res.Next
}

func mergeTwoListsFirstSolution(list1, list2 *l.Node) *l.Node {
	current := &l.Node{Val: 0, Next: nil}
	result := current
	if list1 == nil && list2 == nil {
		return nil
	}

	for list1 != nil || list2 != nil {
		if list2 == nil {
			current.Val = list1.Val
			if list1 != nil && list1.Next != nil {
				current.Next = &l.Node{Val: 0, Next: nil}
				current = current.Next
			}
			list1 = list1.Next
		} else if list1 == nil {
			current.Val = list2.Val
			if list2 != nil && list2.Next != nil {
				current.Next = &l.Node{Val: 0, Next: nil}
				current = current.Next
			}
			list2 = list2.Next
		}
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				current.Val = list1.Val
				current.Next = &l.Node{Val: 0, Next: nil}
				current = current.Next
				list1 = list1.Next
			} else if list2.Val < list1.Val {
				current.Val = list2.Val
				current.Next = &l.Node{Val: 0, Next: nil}
				current = current.Next
				list2 = list2.Next
			} else {
				current.Val = list1.Val
				current.Next = &l.Node{Val: 0, Next: nil}
				current = current.Next
				list1 = list1.Next
			}
		}
	} // end for

	return result
}
