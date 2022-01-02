package is_palindrome

import "math"

// https://leetcode.com/problems/palindrome-number/
func isPalindrome(x int) bool {
	if x < 0 { return false }
	if x < 10 { return true }
	var arr []int
	x = int(math.Abs(float64(x)))
	for x > 0 {
		modResult := x % 10
		arr = append(arr, []int{modResult}...)
		x = x / 10
	}

	for i := 0; i < len(arr)/2; i++ {
		if arr[i] != arr[len(arr)-1-i] {
			return false
		}
	}

	return true
}