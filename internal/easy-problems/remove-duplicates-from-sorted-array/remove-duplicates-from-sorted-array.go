package remove_duplicates_from_sorted_array

func removeDuplicates(nums []int) int {
	nextFreeIndex := 0
	curr := -101
	k := 0
	updateNextFreeIndex := false
	for i := 0; i < len(nums); i++ {
		if curr != nums[i] {
			nums[nextFreeIndex] = nums[i]
			curr = nums[i]
			k++
			updateNextFreeIndex = true
		}
		if updateNextFreeIndex {
			nextFreeIndex++
			updateNextFreeIndex = false
		}
	}

	return k
}