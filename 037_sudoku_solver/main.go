package main

import (
	"fmt"

	"gitea.bojanaxievski.duckdns.org/baxievski/leetcode/037_sudoku_solver/sudokusolver"
)

func main() {
	s := sudokusolver.Sudoku{
		[]int{5, 3, 0, 0, 7, 0, 0, 0, 0},
		[]int{6, 0, 0, 1, 9, 5, 0, 0, 0},
		[]int{0, 9, 8, 0, 0, 0, 0, 6, 0},
		[]int{8, 0, 0, 0, 6, 0, 0, 0, 3},
		[]int{4, 0, 0, 8, 0, 3, 0, 0, 1},
		[]int{7, 0, 0, 0, 2, 0, 0, 0, 6},
		[]int{0, 6, 0, 0, 0, 0, 2, 8, 0},
		[]int{0, 0, 0, 4, 1, 9, 0, 0, 5},
		[]int{0, 0, 0, 0, 8, 0, 0, 7, 9},
	}
	// fmt.Println(s.IsValidSudoku())
	fmt.Println(s.String())
}
