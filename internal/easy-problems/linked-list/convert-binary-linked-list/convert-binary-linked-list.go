package convert_binary_linked_list

// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/

// ListNode: Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

// converts a binary number, represented as a linked list, into a decimal number
func getDecimalValue(head *ListNode) int {
	result := 0
	for ; head != nil; head = head.Next {
		result = (result << 1) + head.Val
	}
	return result
}
