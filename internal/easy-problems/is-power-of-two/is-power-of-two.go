package is_power_of_two

// IsPowerOfTwo https://leetcode.com/problems/power-of-two/
func IsPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	// Logical AND bitwise operation. If both bits are set, the result is 1 (1 x 1 = 1).
	// Otherwise the result is 0 (1 x 0 = 0, 0 x 0 = 0).
	// Any power of 2, in binary, will have a 1 in the first position and 0s in the rest, so if you AND a power of 2
	// with a power of 2 minus 1, you'll get 0.
	return n&(n-1) == 0
}

func IsPowerOfTwoOriginal(n int) bool {
	k := float64(n)
	if n == 1 {
		return true
	}
	for {
		if k == 2 {
			return true
		}
		if k < 2 {
			return false
		}
		k = k / 2
	}
}
