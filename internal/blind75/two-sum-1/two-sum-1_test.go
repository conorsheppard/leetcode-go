package two_sum_1

import (
	"fmt"
	"testing"
)

func Test_TwoSum(t *testing.T) {
	tests := []struct {
		inputArray     []int
		target         int
		expectedResult []int
	}{
		{inputArray: []int{2, 7, 11, 15}, target: 9, expectedResult: []int{0, 1}},
		{inputArray: []int{2, 7, 15, 0}, target: 2, expectedResult: []int{0, 3}},
		{inputArray: []int{2, 7, 0, 15}, target: 2, expectedResult: []int{0, 2}},
		{inputArray: []int{3, 2, 4}, target: 6, expectedResult: []int{1, 2}},
		{inputArray: []int{3, 3}, target: 6, expectedResult: []int{0, 1}},
		{inputArray: []int{-3, 3}, target: 0, expectedResult: []int{0, 1}},
		{
			inputArray:     []int{3, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, 50, 3, 7, 14, 20, -3, 100},
			target:         97,
			expectedResult: []int{240, 241},
		},
	}

	var pass, fail int

	for i, test := range tests {
		fmt.Printf("test %d\n", i)

		result := twoSum(test.inputArray, test.target)

		if fmt.Sprintf("%v\n", result) != fmt.Sprintf("%v\n", test.expectedResult) {
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
