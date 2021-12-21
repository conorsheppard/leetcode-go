package is_power_of_two

import (
	"fmt"
	"testing"
)

// Given an integer n, return true if it is a power of two. Otherwise, return false.
// An integer n is a power of two, if there exists an integer x such that n == 2x.

// Constraints:
// -2^31 <= n <= 2^31 - 1

// https://leetcode.com/problems/power-of-two/
func Test_IsPowerOfTwo(t *testing.T) {
	tests := []struct{
		n int
		isPowerOfTwo bool
	}{
		{
			n: 1,
			isPowerOfTwo: true,
		},
		{
			n: 2,
			isPowerOfTwo: true,
		},
		{
			n: 16,
			isPowerOfTwo: true,
		},
		{
			n: 3,
			isPowerOfTwo: false,
		},
		{
			n: 0,
			isPowerOfTwo: false,
		},
		{
			n: 10,
			isPowerOfTwo: false,
		},
		{
			n: 2147483647,
			isPowerOfTwo: false,
		},
		{
			n: 2147483648,
			isPowerOfTwo: true,
		},
		{
			n: -2147483647,
			isPowerOfTwo: false,
		},
		{
			n: -8,
			isPowerOfTwo: false,
		},
	}

	for i := range tests {
		test := tests[i]
		fmt.Printf("test %d\n", i)
		result := IsPowerOfTwo(test.n)

		if result != test.isPowerOfTwo {
			t.Errorf("test %d: error", i)
		} else {
			t.Logf("test %d: pass", i)
			not := "not "
			if result == true {
				not = ``
			}
			t.Logf("%d is %sa power of 2", test.n, not)
		}
	}
}
