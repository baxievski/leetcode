package sudokusolver

import (
	"fmt"
)

func solveSudoku(board [][]byte) {
	s := make(Sudoku, len(board[0]))
	for i, row := range board {
		s[i] = make([]int, len(board))
		for j, el := range row {
			if el == '.' {
				el = '0'
			}
			s[i][j] = int(el) - int('0')
		}
	}
	s.Solve()
	for i, row := range s {
		for j, el := range row {
			if el == 0 {
				board[i][j] = []byte(".")[0]
				continue
			}
			board[i][j] = []byte(fmt.Sprintf("%d", el))[0]
		}
	}

}

// Sudoku represents a 9x9 sudoku board
type Sudoku [][]int

// Solve tells if it found a solution to the sudoku board
func (s *Sudoku) Solve() bool {
	if s.IsSolved() {
		return true
	}
	for row := range *s {
		for col := range (*s)[row] {
			if (*s)[row][col] == 0 {
				for v := 1; v <= 9; v++ {
					(*s)[row][col] = v
					if s.IsValidSudoku() {
						if s.Solve() {
							return true
						}
						(*s)[row][col] = 0
					}
					(*s)[row][col] = 0
				}
			}
			if (*s)[row][col] == 0 {
				return false
			}
		}
	}
	return true
}

// IsSolved tells if the sudoku board is solved
func (s *Sudoku) IsSolved() bool {
	for _, row := range *s {
		for _, el := range row {
			if el == 0 {
				return false
			}
		}
	}
	return true

}

// IsValidSudoku tells if the structure is a valid sudoku board
func (s *Sudoku) IsValidSudoku() bool {
	if !s.validRows() {
		return false
	}
	if !s.validColumns() {
		return false
	}
	if !s.validSubBoards() {
		return false
	}
	return true
}

func (s *Sudoku) validColumns() bool {
	res := make(Sudoku, len((*s)[0]))
	for i := range res {
		res[i] = make([]int, len(*s))
	}

	for i, r := range *s {
		for j, e := range r {
			res[j][i] = e
		}
	}
	return res.validRows()
}

func (s *Sudoku) validSubBoards() bool {
	res := make(Sudoku, len((*s)[0]))
	for i := range res {
		res[i] = make([]int, len(*s))
	}

	for i, r := range *s {
		for j := range r {
			res[i][j] = (*s)[3*(i/3)+(j/3)][3*(i%3)+(j%3)]
		}
	}
	return res.validRows()
}

func (s *Sudoku) validRows() bool {
	for _, row := range *s {
		seen := map[int]bool{}
		for _, n := range row {
			if n == 0 {
				continue
			}
			if n < 0 {
				return false
			}
			if n > 9 {
				return false
			}
			if seen[n] {
				return false
			}
			seen[n] = true
		}
	}
	return true
}

func (s *Sudoku) String() string {
	o := "\n"
	for _, r := range *s {
		for _, e := range r {
			o += fmt.Sprint(e)
			o += " "
		}
		o += "\n"
	}
	o += "\n"
	return o
}

// Equal tells if 2 sudoku boards are the same
func (s *Sudoku) Equal(o *Sudoku) bool {
	for i := range *s {
		for j := range (*s)[i] {
			if (*s)[j][i] != (*o)[j][i] {
				return false
			}
		}
	}
	return true
}
