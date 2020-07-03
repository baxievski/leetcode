package containerwithmostwater

import "testing"

func TestMaxArea(t *testing.T) {
	for _, test := range tests {
		if maxArea(test.inputs) != test.expected {
			t.Errorf("maxArea(%v) -> %v, expected %v.", test.inputs, maxArea(test.inputs), test.expected)
		}
	}
}

func BenchmarkMaxArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			maxArea(test.inputs)
		}
	}
}
