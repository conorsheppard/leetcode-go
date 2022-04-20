package reverse_string_344

import (
	"fmt"
	"testing"
)

func Test_ReverseString(t *testing.T) {
	tests := []struct {
		input          []byte
		expectedOutput []byte
	}{
		//{
		//	input:          []byte("Hello, 世界 ... this string will be reversed and there's nothing you can do about it!!! Mwuhahahahahahaha 😎"),
		//	expectedOutput: []byte("😎 ahahahahahahahuwM !!!ti tuoba od nac uoy gnihton s'ereht dna desrever eb lliw gnirts siht ... 界世 ,olleH"),
		//},
		{
			input:          []byte("I/O"),
			expectedOutput: []byte("O/I"),
		},
		{
			input:          []byte("ajksdflkjsdbfdsblkjsdalfkjnjskdfbvi slaijbdflaksjdn;jc  sa jsdn;asdjknlfnas ;njf sd"),
			expectedOutput: []byte("ds fjn; sanflnkjdsa;ndsj as  cj;ndjskalfdbjials ivbfdksjnjkfladsjklbsdfbdsjklfdskja"),
		},
		{
			input:          []byte(""),
			expectedOutput: []byte(""),
		},
		{
			input:          nil,
			expectedOutput: nil,
		},
		{
			input:          []byte{},
			expectedOutput: []byte{},
		},
	}

	for i, test := range tests {
		fmt.Printf("test %d\n", i)

		reverseString(test.input)

		if test.input != nil && test.expectedOutput == nil {
			t.Errorf("expected: %s, got: %s\nFAIL\n", test.expectedOutput, test.input)
			continue
		}

		if test.input == nil && test.expectedOutput != nil {
			t.Errorf("expected: %s, got: %s\nFAIL\n", test.expectedOutput, test.input)
			continue
		}

		for i := 0; i < len(test.expectedOutput); i++ {
			if test.input[i] != test.expectedOutput[i] {
				t.Errorf("expected: %s, got: %s\nFAIL\n", test.expectedOutput, test.input)
				break
			}
		}
	}
}
