package linked_list

// https://leetcode.com/explore/learn/card/linked-list/209/singly-linked-list/1290/
type Node struct {
	Val  int
	Next *Node
}

func Constructor() Node {
	return Node{}
}

func (node *Node) Get(index int) int {
	current := node
	for i := 0; i < index; i++ {
		if current.Next == nil {
			return -1
		}
		current = current.Next
	}
	return current.Val
}

func (node *Node) AddAtHead(val int) {
	current := node
	node.Next = current
	node.Val = val
}

func (node *Node) AddAtTail(val int) {
	current := node
	if node.Next == nil {
		node.Next = &Node{Val: val, Next: nil}
		return
	}
	for current.Next != nil {
		current = current.Next
	}
	current.Next = &Node{Val: val, Next: nil}
}

func (node *Node) AddAtIndex(index int, val int) {
	newNode := &Node{Val: val, Next: nil}
	current := node
	for i := 0; i < index-1; i++ {
		if current.Next == nil {
			return
		}
		current = current.Next
	}
	if current.Next == nil {
		current.Next = newNode
	}
	newNode.Next = current.Next
	current.Next = newNode
}

func (node *Node) DeleteAtIndex(index int) {
	if index == 0 {
		node = node.Next
		return
	}
	current := node
	for i := 0; i < index-1; i++ {
		if current.Next == nil {
			return
		}
		current = current.Next
	}
	if current.Next == nil {
		return
	}
	current.Next = current.Next.Next
}
