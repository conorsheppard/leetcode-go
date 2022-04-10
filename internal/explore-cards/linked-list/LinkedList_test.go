package linked_list

import (
	"fmt"
	"testing"
)

// https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list/1290/
func Test_AddInteger(t *testing.T) {
	tests := []struct {
		expectedResult []int
	}{
		{
			//Input: ["MyLinkedList","addAtHead","addAtHead","addAtHead","addAtIndex","deleteAtIndex","addAtHead","addAtTail","get","addAtHead","addAtIndex","addAtHead"]
			//	   [[],[7],[2],[1],[3,0],[2],[6],[4],[4],[4],[5,0],[6]]
			//Output: [null,null,null,null,null,null,null,null,-1,null,null,null]
			//Expected: [null,null,null,null,null,null,null,null,4,null,null,null]
		},
	}

	for i := range tests {
		_ = tests[i]
		fmt.Printf("test %d\n", i)
	}
}
