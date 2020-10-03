package validsudoku

func isValidSudoku(board [][]byte) bool {
	iB := make([][]int, len(board[0]))
	for i, row := range board {
		iB[i] = make([]int, len(board))
		for j, el := range row {
			if el == '.' {
				el = '0'
			}
			iB[i][j] = int(el) - int('0')
		}
	}
	s := Sudoku{Board: iB}
	return s.IsValidSudoku()
}

// Sudoku represents a 9x9 sudoku board
type Sudoku struct {
	Board [][]int
}

// IsValidSudoku tells if the structure is a valid sudoku board
func (s Sudoku) IsValidSudoku() bool {
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

func (s Sudoku) validColumns() bool {
	res := make([][]int, len(s.Board[0]))
	for i := range res {
		res[i] = make([]int, len(s.Board))
	}

	for i, r := range s.Board {
		for j, e := range r {
			res[j][i] = e
		}
	}
	return Sudoku{Board: res}.validRows()
}

func (s Sudoku) validSubBoards() bool {
	res := make([][]int, len(s.Board[0]))
	for i := range res {
		res[i] = make([]int, len(s.Board))
	}

	for i, r := range s.Board {
		for j := range r {
			res[i][j] = s.Board[3*(i/3)+(j/3)][3*(i%3)+(j%3)]
		}
	}
	return Sudoku{Board: res}.validRows()
}

func (s Sudoku) validRows() bool {
	for _, row := range s.Board {
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

func (s Sudoku) String() string {
	o := "\n"
	for _, r := range s.Board {
		for _, e := range r {
			o += string(e)
			o += " "
		}
		o += "\n"
	}
	o += "\n"
	return o
}
