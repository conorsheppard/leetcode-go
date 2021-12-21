package add_array_form_integer

import (
	"fmt"
	"reflect"
	"testing"
)

// https://leetcode.com/problems/add-to-array-form-of-integer/
func Test_AddInteger(t *testing.T) {
	tests := []struct{
		num []int
		k int
		expectedResult []int
	}{
		{
			num: []int{1,2,0,0},
			k: 34,
			expectedResult: []int{1,2,3,4},
		},
		{
			num: []int{2,7,4},
			k: 181,
			expectedResult: []int{4,5,5},
		},
		{
			num: []int{2,1,5},
			k: 806,
			expectedResult: []int{1,0,2,1},
		},
		{
			num: []int{9,9,9,9,9,9,9,9,9,9},
			k: 1,
			expectedResult: []int{1,0,0,0,0,0,0,0,0,0,0},
		},
		{
			num: []int{0},
			k: 0,
			expectedResult: []int{0},
		},
		{
			num: []int{1,0,0,0,0,0},
			k: 100000,
			expectedResult: []int{2,0,0,0,0,0},
		},
	}

	for i := range tests {
		test := tests[i]
		fmt.Printf("test %d\n", i)

		result := addToArrayForm(test.num, test.k)

		if !reflect.DeepEqual(result, test.expectedResult) {
			t.Errorf("test %d: error", i)
			t.Errorf("result: %v != expectedOutput: %v", result, test.expectedResult)
		}
	}
}
