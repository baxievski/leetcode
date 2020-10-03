package validsudoku

import "testing"

func TestIsValidSudoku(t *testing.T) {
	for _, test := range tests {
		s := Sudoku{Board: test.inputs}
		if s.IsValidSudoku() != test.expected {
			t.Errorf("IsValidSudoku() -> %v, expected %v.", s.IsValidSudoku(), test.expected)
		}
	}
}

func BenchmarkIsValidSudoku(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, test := range tests {
			s := Sudoku{Board: test.inputs}
			s.IsValidSudoku()
		}
	}
}
