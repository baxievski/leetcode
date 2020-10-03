package longestcommonprefix

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	for _, test := range tests {
		if longestCommonPrefix(test.inputs) != test.expected {
			t.Errorf("'%v'\nlongestCommonPrefix(%v) -> %v, expected %v.", test.description, test.inputs, longestCommonPrefix(test.inputs), test.expected)
		}
	}
}

func BenchmarkLongestCommonPrefix(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			longestCommonPrefix(test.inputs)
		}
	}
}
