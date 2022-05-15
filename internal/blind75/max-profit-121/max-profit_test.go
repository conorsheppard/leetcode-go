package max_profit_121

import (
	"fmt"
	"testing"
)

func Test_MaxProfit(t *testing.T) {
	tests := []struct {
		inputArray     []int
		expectedResult int
	}{
		{inputArray: []int{7, 1, 5, 3, 6, 4}, expectedResult: 5},
		{inputArray: []int{7, 6, 4, 3, 1}, expectedResult: 0},
	}

	var pass, fail int

	for i, test := range tests {
		fmt.Printf("test %d\n", i)

		result := maxProfit(test.inputArray)

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
