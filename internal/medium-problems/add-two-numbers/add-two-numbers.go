package add_two_numbers

type ListNode struct {
	Val int
	Next *ListNode
}

// AddTwoNumbers problem: https://leetcode.com/problems/add-two-numbers/
func AddTwoNumbers(l1, l2 *ListNode) *ListNode {
	list1 := l1
	list2 := l2
	var carryOver int
	currentNode := &ListNode{Val: 0, Next: nil}
	result := currentNode
	for list1 != nil || list2 != nil {
		var list1Val int
		var list2Val int
		if list1 == nil {
			list1Val = 0
		} else {
			list1Val = list1.Val
		}
		if list2 == nil {
			list2Val = 0
		} else {
			list2Val = list2.Val
		}
		// sum both nodes
		sum := list1Val + list2Val + carryOver
		currentNode.Val = sum % 10
		carryOver = sum / 10
		if list1 != nil {
			list1 = list1.Next
		}
		if list2 != nil {
			list2 = list2.Next
		}
		// adding if so that a new currentNode with a Val of 0 is not created at the end and doesn't add an extra 0
		// to the result ListNode (e.g. 7080)
		if list1 != nil || list2 != nil || carryOver != 0 {
			currentNode.Next = &ListNode{Val: 0, Next: nil}
			// two lines need to be treated as one transaction, otherwise nil pointer will occur as currentNode.Next
			// will be nil
			currentNode = currentNode.Next
		}
		// this is needed for when the loop breaks and the final carryover becomes the most significant digit
		if carryOver != 0 {
			currentNode.Val = carryOver
		}
	}
	return result
}
