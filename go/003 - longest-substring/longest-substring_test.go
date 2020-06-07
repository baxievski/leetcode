package longestsubstring

import "testing"

func TestlengthOfLongestSubstring(t *testing.T) {
	for _, test := range tests {
		if lengthOfLongestSubstring(test.input) != test.expected {
			t.Errorf("lengthOfLongestSubstring(%v) -> %v, expected %v.", test.input, lengthOfLongestSubstring(test.input), test.expected)
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			lengthOfLongestSubstring(test.input)
		}
	}
}
