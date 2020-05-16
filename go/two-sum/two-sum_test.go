package twosum

import "testing"

func TestTwoSum(t *testing.T) {
	for _, test := range tests {
		if actual := twoSum(test.nums, test.target); !equal(actual, test.expected) {
			t.Errorf("twoSum(%v, %v) = %v, expected %v.",
				test.nums, test.target, actual, test.expected)
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			twoSum(test.nums, test.target)
		}
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
