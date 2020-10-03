package sudokusolver

import (
	"testing"
)

func TestSolve(t *testing.T) {
	for _, test := range tests {
		s := test.inputs
		solved := s.Solve()
		if solved != test.solved {
			t.Errorf("Solve() -> %v, expected %v.", solved, test.solved)
		}
		if !s.Equal(&test.result) {
			t.Errorf("Solve() -> %s, expected %s.", s.String(), test.result.String())
		}
	}
}

func BenchmarkIsValidSudoku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			s := test.inputs
			s.IsValidSudoku()
		}
	}
}
