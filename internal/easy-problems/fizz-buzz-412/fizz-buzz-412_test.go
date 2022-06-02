package fizz_buzz_412

import (
	"testing"
)

func Test_FizzBuzz(t *testing.T) {
	tests := []struct {
		input          int
		expectedResult []string
	}{
		{
			input:          1,
			expectedResult: []string{"1"},
		},
		{
			input:          3,
			expectedResult: []string{"1", "2", "Fizz"},
		},
		{
			input:          5,
			expectedResult: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			input:          15,
			expectedResult: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	var pass, fail = len(tests), 0

	for i, test := range tests {
		t.Logf("test %d", i)

		result := fizzBuzz(test.input)

		for j := 0; j < len(test.expectedResult); j++ {
			if len(test.expectedResult) != len(result) {
				t.Errorf("FAIL: arrays of unequal length: expected: %d, got: %d\n", len(test.expectedResult), len(result))
				failTest(t, &fail, &pass, "", test.expectedResult[j], result[j])
			}
			if result[j] != test.expectedResult[j] {
				failTest(t, &fail, &pass, "", test.expectedResult[j], result[j])
				break
			}
		}
	}

	t.Logf("\n")
	t.Logf("Tests ran: %d\n", len(tests))
	t.Logf("Tests passed: %d\n", pass)
	t.Logf("Tests failed: %d\n", fail)
}

func failTest(t *testing.T, fail, pass *int, message string, s ...string) {
	*fail++
	*pass--
	t.Errorf("FAIL: expected %s, got %s\n", s[0], s[1])
}
