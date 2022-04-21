package is_palindrome_linked_list_234

import (
	l "leetcode/internal/easy-problems/linked-list"
)

// Time: O(n), Space: O(1)
func isPalindrome(head *l.Node) bool {
	listLen, cur := 0, head
	for ; cur != nil; listLen, cur = listLen+1, cur.Next {
	}

	for i := 0; i < listLen/2; head, i = head.Next, i+1 {
		cur = &l.Node{Val: head.Val, Next: cur}
	}

	if listLen%2 == 1 {
		cur = &l.Node{Val: head.Val, Next: cur}
	}

	for ; head != nil; head, cur = head.Next, cur.Next {
		if cur.Val != head.Val {
			return false
		}
	}

	return true
}

// Time: O(n), Space: O(1)
func isPalindromeOXR(head *l.Node) bool {
	lngth := 0
	for p := head; p != nil; p = p.Next {
		lngth++
	}

	if lngth == 0 || lngth == 1 {
		return true
	}

	c := make([]int, 10)
	for i, p := 0, head; p != nil; i, p = i+1, p.Next {
		if i >= lngth/2 {
			c[p.Val] ^= i
		} else {
			c[p.Val] ^= lngth - 1 - i
		}
	}

	nonZeros := 0
	for _, v := range c {
		if v != 0 {
			nonZeros += 1
		}
	}

	return nonZeros == 0 || nonZeros == 1
}

// Time: O(n), Space: O(1)
func isPalindromeXOR2(head *l.Node) bool {
	ln1, ln2, xr := 0, 0, 0
	for node := head; node != nil; ln1, node = ln1+1, node.Next {
	}
	for node := head; node != nil; node = node.Next {
		if ln2 >= ln1/2+ln1%2 {
			xr ^= 11*(ln1-ln2-1) + node.Val
		} else if ln2 < ln1/2 {
			xr ^= 11*(ln2) + node.Val
		}
		ln2 += 1
	}

	return xr == 0
}
