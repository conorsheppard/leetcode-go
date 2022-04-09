package add_two_numbers

import (
	"fmt"
	"strconv"
)

// AddTwoNumbers problem: https://leetcode.com/problems/add-two-numbers/
func AddTwoNumbersFirstSolution(l1, l2 *ListNode) *ListNode {
	list1Total := extractNumber(l1)
	list2Total := extractNumber(l2)

	combinedNums := list1Total + list2Total

	result := addNextNodes(strconv.Itoa(combinedNums))
	return result
}

func extractNumber(list *ListNode) int {
	list1NumStr := ""
	for {
		if list.Next != nil {
			list1NumStr = strconv.Itoa(list.Val) + list1NumStr
		} else {
			list1NumStr = strconv.Itoa(list.Val) + list1NumStr
			break
		}
		list = list.Next
	}

	list1NumInt, err := strconv.Atoi(list1NumStr)

	if err != nil {
		fmt.Println(err.Error())
	}

	return list1NumInt
}

func addNextNodes(combinedNums string) *ListNode {
	valAsStr := string(combinedNums[len(combinedNums)-1])
	val, _ := strconv.Atoi(valAsStr)

	if len(combinedNums)-1 == 0 {
		return &ListNode{
			Val:  val,
			Next: nil,
		}
	}

	combinedNums = combinedNums[:len(combinedNums)-1]
	return &ListNode{
		Val:  val,
		Next: addNextNodes(combinedNums),
	}
}

func acceptedSolution(l1, l2 *ListNode) *ListNode {
	current := new(ListNode)
	result := current
	sum := 0

	for l1 != nil || l2 != nil {
		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{
			sum % 10,
			nil,
		}

		current = current.Next
		sum = sum / 10
	}

	if sum > 0 {
		current.Next = &ListNode{sum, nil}
	}

	return result.Next
}