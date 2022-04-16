package delete_node_from_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
	// *node = *(node.Next)
}
