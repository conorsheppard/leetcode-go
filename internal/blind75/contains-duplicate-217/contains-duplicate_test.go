package contains_duplicate_217

import (
	"fmt"
	"testing"
)

func Test_ContainsDuplicate(t *testing.T) {
	tests := []struct {
		inputArray     []int
		expectedResult bool
	}{
		{inputArray: []int{1, 2, 3, 1}, expectedResult: true},
		{inputArray: []int{1, 2, 3, 4}, expectedResult: false},
		{inputArray: []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, expectedResult: true},
		{inputArray: []int{1}, expectedResult: false},
	}

	var pass, fail int

	for i, test := range tests {
		fmt.Printf("test %d\n", i)

		result := containsDuplicate(test.inputArray)

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
