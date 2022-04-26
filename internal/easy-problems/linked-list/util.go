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

func CheckListEquality(result, expectedList *Node) (bool, string) {
	for result != nil || expectedList != nil {
		if result == nil && expectedList != nil {
			return false, "not equal, result is nil, test.expectedList not nil\n"
		} else if result != nil && expectedList == nil {
			return false, "not equal, result is not nil, test.expectedList is nil\n"
		} else if result == nil && expectedList == nil {
			return true, "lists are both nil. PASS.\n"
		}
		if result.Val != expectedList.Val {
			return false, "expected: " + GetLinkedListAsString(expectedList) + ", got: " + GetLinkedListAsString(result) + "\n"
		}
		result = result.Next
		expectedList = expectedList.Next
	}

	return true, "PASS\n"
}
