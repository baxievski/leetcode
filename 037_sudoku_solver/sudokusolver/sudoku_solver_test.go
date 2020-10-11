package sudokusolver

import (
	"testing"
)

func TestSolve(t *testing.T) {
	for _, test := range tests {
		s := test.inputs
		// fmt.Printf("%s", s.String())
		// for i, p := range s.AllPossibilities() {
		// 	if len(p) == 1 {
		// 		fmt.Printf("%v: %v\n", i, p)
		// 	}
		// }
		solved := s.Solve()
		// fmt.Printf("solved:%v", s.String())
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
			s.IsValid()
		}
	}
}
