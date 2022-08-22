package plus_minus_zero

import "fmt"

// https://www.hackerrank.com/challenges/plus-minus/problem
func plusMinusZero(arr []int32) {
	// Write your code here
	var plusCount, minusCount, zeroCount float32

	for _, num := range arr {
		if num > 0 {
			plusCount++
		} else if num < 0 {
			minusCount++
		} else {
			zeroCount++
		}
	}

	plusRatio := plusCount / float32(len(arr))
	minusRatio := minusCount / float32(len(arr))
	zeroRatio := zeroCount / float32(len(arr))

	fmt.Printf("%.6f\n", plusRatio)
	fmt.Printf("%.6f\n", minusRatio)
	fmt.Printf("%.6f\n", zeroRatio)
}
