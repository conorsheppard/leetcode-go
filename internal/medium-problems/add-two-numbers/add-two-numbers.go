package add_two_numbers

type ListNode struct {
	Val  int
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

func AddTwoNumbers2(l1, l2 *ListNode) *ListNode {
	sum := 0
	carryOver := 0
	current := &ListNode{Val: 0, Next: nil}
	result := current // &ListNode{Val: 0, Next: &ListNode{Val: 8, Next: &ListNode{Val: 0, Next: nil}}}

	for l1 != nil || l2 != nil {
		if l1 == nil {
			sum = 0 + l2.Val + carryOver
		} else if l2 == nil {
			sum = l1.Val + 0 + carryOver
		} else {
			sum = l1.Val + l2.Val + carryOver
		}

		if sum >= 10 {
			carryOver = 1
			sum = sum - 10
		} else {
			carryOver = 0
		}

		current.Next = &ListNode{Val: sum, Next: nil}
		current = current.Next

		if carryOver == 0 && l1.Next == nil {
			current.Next = l2.Next
			break
		} else if carryOver == 0 && l2.Next == nil {
			current.Next = l1.Next
			break
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
		if carryOver == 1 && l1 == nil && l2 == nil {
			current = &ListNode{Val: carryOver, Next: nil}
		}
	}

	return result.Next
}
