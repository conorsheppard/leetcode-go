package linked_list

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetLinkedListAsString(l1 *ListNode) string {
	listStr := "["
	for ; l1 != nil; l1 = l1.Next {
		listStr = fmt.Sprintf("%s%d", listStr, l1.Val)
		if l1.Next != nil {
			listStr = fmt.Sprintf("%s, ", listStr)
		} else {
			listStr = fmt.Sprintf("%s]\n", listStr)
		}
	}
	return listStr
}
