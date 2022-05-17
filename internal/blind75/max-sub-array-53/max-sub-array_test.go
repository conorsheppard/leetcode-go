package max_sub_array_53

import (
	"fmt"
	"testing"
)

func Test_MaxSubArray(t *testing.T) {
	tests := []struct {
		inputArray     []int
		expectedResult int
	}{
		{inputArray: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, expectedResult: 6},
		{inputArray: []int{1}, expectedResult: 1},
		{inputArray: []int{5, 4, -1, 7, 8}, expectedResult: 23},
	}

	var pass, fail int

	for i, test := range tests {
		fmt.Printf("test %d\n", i)

		result := maxSubArray(test.inputArray)

		if result != test.expectedResult {
			t.Errorf("FAIL: test %d\n", i)
			t.Errorf("Expected: %v, Got: %v\n", test.expectedResult, result)
			fail++
		} else {
			pass++
		}
	}
	fmt.Printf("Tests ran: %d\n", len(tests))
	fmt.Printf("Tests passed: %d\n", pass)
	fmt.Printf("Tests failed: %d\n", fail)
}
