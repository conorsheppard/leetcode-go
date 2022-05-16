package product_except_self_238

func productExceptSelf(nums []int) (arr []int) {
	var zeroIndices, totalProduct = map[int]bool{}, 1

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			zeroIndices[i] = true
		} else {
			totalProduct = totalProduct * nums[i]
		}
	}

	for j := 0; j < len(nums); j++ {
		if len(zeroIndices) > 0 {
			if zeroIndices[j] && len(zeroIndices) == 1 {
				arr = append(arr, totalProduct)
			} else {
				arr = append(arr, 0)
			}
		} else {
			arr = append(arr, totalProduct/nums[j])
		}
	}

	return
}
