package stack

import (
	"errors"
	"fmt"
	"testing"
)

func Test_MinStack(t *testing.T) {

	tests := []struct {
		runTest func() error
	}{
		{
			runTest: func() error {
				// push, pop return the top element and check the result
				ms := Constructor()
				ms.Push(1)
				ms.Pop()
				result := ms.Top()
				if result != MaxStackSize {
					return errors.New("error: .Top() should return MaxStackSize (" + fmt.Sprint(MaxStackSize) + ")")
				}
				return nil
			},
		},
		{
			runTest: func() error {
				ms := Constructor()
				ms.Push(3)
				ms.Push(2)
				ms.Push(1)
				min := ms.GetMin()
				if min != 1 {
					return errors.New("error: min should be 1")
				}
				return nil
			},
		},
		{
			runTest: func() error {
				ms := Constructor()
				min := ms.GetMin()
				if min != MaxStackSize {
					return errors.New("error: min should be " + fmt.Sprint(MaxStackSize))
				}
				return nil
			},
		},
		{
			runTest: func() error {
				ms := Constructor()
				ms.Pop()
				return nil
			},
		},
		{
			runTest: func() error {
				ms := Constructor()
				for i := 0; i < (MaxStackSize); i++ {
					ms.Push(i)
				}
				return nil
			},
		},
	}

	var pass, fail = len(tests), 0

	for i, test := range tests {
		t.Logf("test %d\n", i)

		err := test.runTest()

		if err != nil {
			failTest(t, &pass, &fail, err.Error())
		}
	}

	t.Logf("\n")
	t.Logf("Tests ran: %d", len(tests))
	t.Logf("Tests passed: %d", pass)
	t.Logf("Tests failed: %d", fail)
}

func failTest(t *testing.T, pass, fail *int, message string) {
	*fail++
	*pass--
	t.Errorf(message)
}
