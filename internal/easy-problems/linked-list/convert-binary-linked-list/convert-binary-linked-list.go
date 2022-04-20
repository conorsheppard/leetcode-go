package convert_binary_linked_list

import l "leetcode/internal/easy-problems/linked-list"

// https://leetcode.com/problems/convert-binary-number-in-a-linked-list-to-integer/
// converts a binary number, represented as a linked list, into a decimal number
func getDecimalValue(head *l.Node) int {
	result := 0
	for ; head != nil; head = head.Next {
		result = (result << 1) + head.Val
	}
	return result
}
