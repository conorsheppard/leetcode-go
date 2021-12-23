package merge_two_sorted_lists

// https://leetcode.com/problems/merge-two-sorted-lists/

// ListNode Definition for singly-linked list.
type ListNode struct {
    Val int
    Next *ListNode
}

// list1 [1,2,3,9]
// list1 [7,8,9]
// &ListNode{Val: 7, Next: nil}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	curr := &ListNode{}
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

func mergeTwoListsFirstSolution(list1, list2 *ListNode) *ListNode {
	current := &ListNode{Val: 0, Next: nil}
	result := current
	if list1 == nil && list2 == nil {
		return nil
	}

	for list1 != nil || list2 != nil {
		if list2 == nil {
			current.Val = list1.Val
			if list1 != nil && list1.Next != nil {
				current.Next = &ListNode{Val: 0, Next: nil}
				current = current.Next
			}
			list1 = list1.Next
		} else if list1 == nil {
			current.Val = list2.Val
			if list2 != nil && list2.Next != nil {
				current.Next = &ListNode{Val: 0, Next: nil}
				current = current.Next
			}
			list2 = list2.Next
		}
		if list1 != nil && list2 != nil {
			if list1.Val < list2.Val {
				current.Val = list1.Val
				current.Next = &ListNode{Val: 0, Next: nil}
				current = current.Next
				list1 = list1.Next
			} else if list2.Val < list1.Val {
				current.Val = list2.Val
				current.Next = &ListNode{Val: 0, Next: nil}
				current = current.Next
				list2 = list2.Next
			} else {
				current.Val = list1.Val
				current.Next = &ListNode{Val: 0, Next: nil}
				current = current.Next
				list1 = list1.Next
			}
		}
	} // end for

	return result
}