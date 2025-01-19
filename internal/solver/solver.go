package solver

import (
	"fmt"
	"sudoku-solver/internal/board"
)

// Solver represents a Sudoku solving algorithm
type Solver interface {
	Solve(b *board.Board) (bool, error)
}

// BacktrackingSolver implements the Solver interface using a backtracking algorithm
type BacktrackingSolver struct{}

func NewBacktrackingSolver() *BacktrackingSolver {
	return &BacktrackingSolver{}
}

var itter = 0

func (s *BacktrackingSolver) Solve(b *board.Board) (bool, error) {
	itter++
	fmt.Println("Itteration: %d", itter)
	// Find empty cell
	row, col, found := findEmptyCell(b)
	if !found {
		// No empty cells means we've solved the puzzle
		return true, nil
	}

	// Try digits 1-9
	for num := 1; num <= 9; num++ {
		// Check if number is valid in this position
		valid, err := b.IsValidMove(row, col, num)
		if err != nil {
			return false, err
		}

		if valid {
			// Try this number
			err := b.SetValue(row, col, num)
			if err != nil {
				return false, err
			}

			// Recursively solve rest of the board
			solved, err := s.Solve(b)
			if err != nil {
				return false, err
			}
			if solved {
				return true, nil
			}

			// If not solved, backtrack by clearing the cell
			err = b.SetValue(row, col, 0)
			if err != nil {
				return false, err
			}
		}
	}

	// No valid solution found
	return false, nil
}

// findEmptyCell finds the next empty cell in the board
func findEmptyCell(b *board.Board) (row, col int, found bool) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			val, err := b.GetValue(i, j)
			if err != nil {
				continue
			}
			if val == 0 {
				return i, j, true
			}
		}
	}
	return -1, -1, false
}
