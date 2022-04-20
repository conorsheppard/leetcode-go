package linked_list

import (
	"fmt"
)

type Node struct {
	Val  int
	Next *Node
}

func GetLinkedListAsString(l1 *Node) string {
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
