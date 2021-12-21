package add_array_form_integer

// AddInteger https://leetcode.com/problems/add-to-array-form-of-integer/
func addToArrayForm(num []int, k int) []int {
	var result []int
	carryOver := 0
	var leastSigDigitNum int
	var leastSigDigitK int

	for {
		if len(num) != 0 {
			// return an array with 1 element and index into it at 0
			leastSigDigitNum = num[len(num)-1:][0]
			// pop off last element in array
			num = num[:len(num)-1]
		} else {
			leastSigDigitNum = 0
		}
		// k % 10 != 0 can't be used because if k is 80 then k % 10 != 0 will be false
		if k != 0 {
			leastSigDigitK = k % 10
			k = k / 10
		} else {
			leastSigDigitK = 0
		}
		sum := leastSigDigitNum + leastSigDigitK + carryOver
		carryOver = sum / 10
		value := []int{sum % 10}
		result = append(value, result...)
		if k == 0 && len(num) == 0 && carryOver != 1 {
			break
		}
	}
	return result
}
