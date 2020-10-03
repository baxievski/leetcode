package validparentheses

import "testing"

func TestIsValid(t *testing.T) {
	for _, test := range tests {
		if isValid(test.input) != test.expected {
			t.Errorf("isValid(%v) -> %v, expected %v.", test.input, isValid(test.input), test.expected)
		}
	}
}

func BenchmarkLengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			isValid(test.input)
		}
	}
}
