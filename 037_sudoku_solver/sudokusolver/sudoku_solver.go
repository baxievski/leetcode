package sudokusolver

import (
	"fmt"
	"strings"
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
	if !s.IsValid() {
		return false
	}
	for row := range *s {
		for col := range (*s)[row] {
			if (*s)[row][col] == 0 {
				for _, v := range s.moves(row, col) {
					(*s)[row][col] = v
					if s.Solve() {
						return true
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

// IsValid tells if the structure is a valid sudoku board
func (s *Sudoku) IsValid() bool {
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
		seen := map[int]bool{}
		for j := range (*s)[i] {
			el := (*s)[j][i]
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
func (s *Sudoku) moves(i int, j int) []int {
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
		// seen in row
		seen[(*s)[i][k]] = true
		// seen in column
		seen[(*s)[k][j]] = true
		// seen in subboard
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

func (s *Sudoku) String() string {
	sb := strings.Builder{}
	sb.WriteString("\n")
	for row, r := range *s {
		if row%3 == 0 {
			sb.WriteString("+------+------+------+\n")
		}
		for col, e := range r {
			if col%3 == 0 {
				sb.WriteString("|")
			}
			sb.WriteString(fmt.Sprintf("%v ", e))
		}
		sb.WriteString("|\n")
	}
	sb.WriteString("+------+------+------+\n")
	return sb.String()
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
