package is_one_bit_char

import (
	"fmt"
	"testing"
)

/**
	0 =< bits[i] <= 1
	1 =< len(bits) <= 1000
	bits[len(bits)-1] == 0
	1st char == 0, 2nd char == 10 || 11
	return true if the last character must be a 1-bit character
	bits slice always ends with a 0 bit
	if you fall on a 1 it has to be a 2-bit character (10 || 11)
	if you skip over the following character when you fall on a 1, then ever 0 has to be a 1-bit character
**/

//https://leetcode.com/problems/1-bit-and-2-bit-characters/
func Test_IsOneBitCharacter(t *testing.T) {
	tests := []struct {
		bits           []int
		expectedOutput bool
	}{
		{[]int{0}, true},

		{[]int{0, 0}, true},
		{[]int{1, 0}, false},

		{[]int{0, 1, 0}, false},
		{[]int{1, 0, 0}, true},
		{[]int{1, 1, 0}, true},

		{[]int{0, 1, 0, 0}, true},
		{[]int{0, 1, 1, 0}, true},
		{[]int{1, 0, 0, 0}, true},
		{[]int{1, 0, 1, 0}, false},
		{[]int{1, 1, 0, 0}, true},
		{[]int{1, 1, 1, 0}, false},
	}

	for i := range tests {
		test := tests[i]
		fmt.Printf("test %d\n", i)

		isOneBitChar := isOneBitCharacter(test.bits)

		if isOneBitChar != test.expectedOutput {
			t.Errorf("expected %t, got %t", test.expectedOutput, isOneBitChar)
		} else {
			fmt.Println("PASS!!")
		}

	}
}
