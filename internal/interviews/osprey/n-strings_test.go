package osprey

import (
	"fmt"
	"testing"
)

/*
take = (["apple", "pear", "lemon", "orange"], 2)
# ["apple", "pear"]
*/

/*
take = (["apple", "pear", "lemon", "orange"], 5)
# ["apple", "pear", "lemon", "orange", "apple"]
*/

func TestNStrings(t *testing.T) {
	tests := []struct {
		inputStr       []string
		inputNum       int
		expectedResult []string
	}{
		{
			inputStr:       []string{},
			inputNum:       2,
			expectedResult: []string{},
		},
		{
			inputStr:       []string{"apple", "pear", "lemon", "orange"},
			inputNum:       2,
			expectedResult: []string{"apple", "pear"},
		},
		{
			inputStr:       []string{"apple", "pear", "lemon", "orange"},
			inputNum:       5,
			expectedResult: []string{"apple", "pear", "lemon", "orange", "apple"},
		},
		{
			inputStr:       []string{"apple", "pear", "lemon", "orange"},
			inputNum:       10,
			expectedResult: []string{"apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear"},
		},
		{
			inputStr:       []string{"apple", "pear", "lemon", "orange"},
			inputNum:       20,
			expectedResult: []string{"apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange"},
		},
		{
			inputStr:       []string{"apple", "pear", "lemon", "orange"},
			inputNum:       80,
			expectedResult: []string{"apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange"},
		},
		{
			inputStr:       []string{"apple", "pear", "lemon", "orange"},
			inputNum:       100,
			expectedResult: []string{"apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange", "apple", "pear", "lemon", "orange"},
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d\n", i)
		result := nStrings(test.inputStr, test.inputNum)

		if fmt.Sprintf("%v", result) != fmt.Sprintf("%v", test.expectedResult) {
			t.Errorf("FAIL: expected: %v, got: %v", fmt.Sprintf("%v", test.expectedResult), fmt.Sprintf("%v", result))
		} else {
			fmt.Printf("\tPASS\n")
		}
	}
}
