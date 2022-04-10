package main

import (
	"fmt"
	linked_list "leetcode/internal/explore-cards/linked-list"
)

func main() {
	linkedList := linked_list.Constructor()
	linkedList.AddAtTail(3)

	fmt.Println(linkedList.Get(0))
}
