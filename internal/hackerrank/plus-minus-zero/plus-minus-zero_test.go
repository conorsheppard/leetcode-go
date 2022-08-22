package plus_minus_zero

import (
	"fmt"
	"testing"
)

func Test_PlusMinusZero(t *testing.T) {
	tests := []struct {
		input []int32
	}{
		{
			input: []int32{-4, 3, -9, 0, 4, 1},
		},
	}

	for i, test := range tests {
		fmt.Printf("test: %d", i)
		plusMinusZero(test.input)
	}
}
