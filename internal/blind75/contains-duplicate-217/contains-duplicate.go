package contains_duplicate_217

func containsDuplicate(nums []int) (result bool) {
	var numsMap = make(map[int]int)
	for _, v := range nums {
		if _, p := numsMap[v]; p {
			return true
		} else {
			numsMap[v] = 1
		}
	}
	return
}
