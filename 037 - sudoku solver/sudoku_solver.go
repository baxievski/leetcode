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
				for _, v := range s.posibilities(row, col) {
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
	for i := range *s {
		seenInColumn := map[int]bool{}
		for j := range (*s)[i] {
			elInColumn := (*s)[j][i]
			if elInColumn == 0 {
				continue
			}
			if elInColumn < 0 {
				return false
			}
			if elInColumn > 9 {
				return false
			}
			if seenInColumn[elInColumn] {
				return false
			}
			seenInColumn[elInColumn] = true
		}
	}
	return true
}

func (s *Sudoku) validRows() bool {
	for i := range *s {
		seen := map[int]bool{}
		for j := range (*s)[i] {
			el := (*s)[i][j]
			if el == 0 {
				continue
			}
			if el < 0 {
				return false
			}
			if el > 9 {
				return false
			}
			if seen[el] {
				return false
			}
			seen[el] = true
		}
	}
	return true
}

func (s *Sudoku) validSubBoards() bool {
	for i := range *s {
		seen := map[int]bool{}
		for j := range (*s)[i] {
			el := (*s)[3*(i/3)+(j/3)][3*(i%3)+(j%3)]
			if el == 0 {
				continue
			}
			if el < 0 {
				return false
			}
			if el > 9 {
				return false
			}
			if seen[el] {
				return false
			}
			seen[el] = true
		}
	}
	return true
}
func (s *Sudoku) posibilities(i int, j int) []int {
	seen := map[int]bool{
		1: false,
		2: false,
		3: false,
		4: false,
		5: false,
		6: false,
		7: false,
		8: false,
		9: false,
	}
	for k := range *s {
		// remove used in row
		seen[(*s)[i][k]] = true
		// remove used in column
		seen[(*s)[k][j]] = true
		// remove used in subboard
		seen[(*s)[3*(i/3)+(k/3)][(k%3)+3*(j/3)]] = true
	}
	p := []int{}
	for k, v := range seen {
		if !v {
			p = append(p, k)
		}
	}
	return p
}

// Row gives all the elements for the ith row
func (s *Sudoku) Row(i int, _ int) []int {
	r := make([]int, len(*s))
	for k := range *s {
		r[k] = (*s)[i][k]
	}
	return r
}

// Column gives all the elements for the jth column
func (s *Sudoku) Column(_ int, j int) []int {
	r := make([]int, len(*s))
	for k := range *s {
		r[k] = (*s)[k][j]
	}
	return r
}

// SubBoard gives all the elements for the subboard in which i, j is located
func (s *Sudoku) SubBoard(i int, j int) []int {
	r := make([]int, len(*s))
	for k := range *s {
		r[k] = (*s)[3*(i/3)+(k/3)][(k%3)+3*(j/3)]
	}
	return r
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
