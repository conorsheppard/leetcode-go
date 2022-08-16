package delete_node_from_linked_list

import l "leetcode-go/internal/easy-problems/linked-list"

func deleteNode(node *l.Node) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
	// *node = *(node.Next)
}
