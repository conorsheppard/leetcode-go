package roman_to_int

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/roman-to-integer/
// Constraints:
// 1 <= input <= 15
// s contains only the characters ('I', 'V', 'X', 'L', 'C', 'D', 'M').
// It is guaranteed that s is a valid roman numeral in the range [1, 3999].
func Test_RomanToInt(t *testing.T) {
	tests := []struct{
		input string
		expectedOutput int
	}{
		{`I`, 1},
		{`II`, 2},
		{`III`, 3},
		{`IV`, 4},
		{`V`, 5},
		{`IX`, 9},
		{`X`, 10},
		{`XXVII`, 27},
		{`L`, 50},
		{`LVIII`, 58},
		{`XC`, 90},
		{`C`, 100},
		{`D`, 500},
		{`CM`, 900},
		{`M`, 1000},
		{`MCMXCIV`, 1994},
		{`MMMCMXCIX`, 3999},
	}

	for i := range tests {
		test := tests[i]
		fmt.Printf("test %d\n", i)

		output := romanToInt(test.input)

		if output != test.expectedOutput {
			t.Errorf("output %d is not equal to %d", output, test.expectedOutput)
		}
	}
}
