package is_palindrome_string

import (
	llq "github.com/emirpasic/gods/queues/linkedlistqueue"
	lls "github.com/emirpasic/gods/stacks/linkedliststack"
	"strings"
)

func isPalindrome(s string) bool {
	chars := []rune(strings.ToLower(s))
	queue := llq.New()
	stack := lls.New()
	for _, char := range chars {
		queue.Enqueue(char)
		stack.Push(char)
	}
	for i := 0; i <= queue.Size(); i++ {
		qChar, _ := queue.Dequeue()
		sChar, _ := stack.Pop()
		if qChar != sChar {
			return false
		}
	}
	return true
}
