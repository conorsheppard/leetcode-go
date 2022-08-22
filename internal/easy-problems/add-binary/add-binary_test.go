package add_binary

import (
	"fmt"
	"testing"
)

// Given two binary strings a and b, return their sum as a binary string.

// Constraints:
// 1 <= a.length, b.length <= 10^4
// a and b consist only of '0' or '1' characters.
// Each string does not contain leading zeros except for the number zero itself.

func Test_AddBinary(t *testing.T) {
	tests := []struct {
		inputA         string
		inputB         string
		expectedOutput string
	}{
		{
			inputA:         `11`,
			inputB:         `1`,
			expectedOutput: `100`,
		},
		{
			inputA:         `1010`,
			inputB:         `1011`,
			expectedOutput: `10101`,
		},
		{
			inputA:         `0`,
			inputB:         `0`,
			expectedOutput: `0`,
		},
	}

	for i := range tests {
		test := tests[i]
		fmt.Printf("test %d\n", i)
		result := addBinary(test.inputA, test.inputB)

		if result != test.expectedOutput {
			t.Errorf("test %d: error", i)
			t.Errorf("result: %s != expectedOutput: %s", result, test.expectedOutput)
		}
	}
}
