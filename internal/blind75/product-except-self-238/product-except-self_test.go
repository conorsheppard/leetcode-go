package product_except_self_238

import (
	"fmt"
	"testing"
)

func Test_ProductExceptSelf(t *testing.T) {
	tests := []struct {
		inputArray     []int
		expectedResult []int
	}{
		{inputArray: []int{1, 2, 3, 4}, expectedResult: []int{24, 12, 8, 6}},
		{inputArray: []int{-1, 1, 0, -3, 3}, expectedResult: []int{0, 0, 9, 0, 0}},
		{inputArray: []int{-1, 1, 0, -3, 3, 0}, expectedResult: []int{0, 0, 0, 0, 0, 0}},
	}

	var pass, fail int

	for i, test := range tests {
		fmt.Printf("test %d\n", i)

		result := productExceptSelf(test.inputArray)

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
