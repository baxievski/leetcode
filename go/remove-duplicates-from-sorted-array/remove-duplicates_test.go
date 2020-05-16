package removeduplicates

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	for _, test := range tests {
		cpy := make([]int, len(test.nums))
		copy(cpy, test.nums)
		if actualLen := removeDuplicates(test.nums); !equal(test.nums[:actualLen], test.expected) {
			t.Errorf("removeDuplicates(%v) -> %v, expected %v.", cpy, test.nums[:actualLen], test.expected)
		}
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			removeDuplicates(test.nums)
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
