package merge_k_lists_23

import (
	l "leetcode/internal/easy-problems/linked-list"
)

func mergeKLists(lists []*l.Node) (result *l.Node) {
	if len(lists) == 0 {
		return
	}
	var valMap = make(map[int]int)
	// if !isPresent add a new entry in the array
	// 1, 4, 5, 3, 2, 6
	for i := 0; i < len(lists); i++ {
		j := 0
		for temp := lists[i]; temp != nil; temp, j = temp.Next, j+1 {
			valMap[temp.Val] = valMap[temp.Val] + 1
		}
	}
	result = &l.Node{Val: 0, Next: nil}
	list := result
	for v := range valMap {
		for i := 0; i < valMap[v]; i++ {
			result.Next = &l.Node{Val: v, Next: nil}
			result = result.Next
		}
	}

	return list.Next
}
