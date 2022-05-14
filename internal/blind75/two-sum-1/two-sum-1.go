package two_sum_1

func twoSum0n2(nums []int, target int) (result []int) {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return
}

func twoSum(nums []int, target int) (result []int) {
	// map: [ number in the list | pointer to the index ]
	var hm = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, isPresent := hm[target-nums[i]]; isPresent {
			return []int{hm[target-nums[i]], i}
		}
		hm[nums[i]] = i
	}
	return
}
