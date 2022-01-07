package remove_duplicates_from_sorted_array

import (
	"fmt"
	"testing"
)

// https://leetcode.com/problems/remove-duplicates-from-sorted-array/
func Test_RemoveDuplicates(t *testing.T) {
	tests := []struct{
		inputArray []int
		expectedNums []int
	}{
		{
			[]int{1,1,2},
			[]int{1,2,},
		},
		{
			[]int{0,0,1,1,1,2,2,3,3,4},
			[]int{0,1,2,3,4},
		},
		{
			[]int{-10,-4,-4,0,0,1,1,1,2,2,3,3,4},
			[]int{-10,-4,0,1,2,3,4},
		},
	}

	for i := range tests {
		test := tests[i]
		fmt.Printf("test %d\n", i)
		nums := test.inputArray
		k := removeDuplicates(nums)

		if k != len(test.expectedNums) {
			t.Errorf("expected %d but k was length %d", len(test.expectedNums), k)
		}
		for j := 0; j < k; j++ {
			if nums[j] != test.expectedNums[j] {
				t.Errorf("expected %d but nums at index %d was %d", test.expectedNums[j], j, nums[j])
			}
		}
	}
}
