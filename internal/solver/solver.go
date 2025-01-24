package solver

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

func (s *BacktrackingSolver) Solve(b *board.Board, depth int, itter *int) (bool, error) {
	*itter++
	fmt.Printf("Running at Depth: %d - Itter: %d \n", depth, *itter)
	// Find empty cell
	row, col, found := findEmptyCell(b)
	if !found {
		// TODO: Check and confirm is puzzle solved sucesfully.
		return true, nil
	}

	// ForEach Number
	// Check to see if Valid
	// If Valid, Place Number
	// Pass Board forwards for next Itteration.
	// Try digits 1-9
	for num := 1; num <= 9; num++ {
		valid, err := b.IsValidMove(row, col, num)
		if err != nil {
			return false, err
		}

		if !valid {
			continue
		}
		err = b.SetValue(row, col, num)
		if err != nil {
			return false, err
		}

		// Lets Watch it go BURR
		clearScreen()
		b.Print()

		solved, err := s.Solve(b, depth+1, itter)
		if err != nil {
			return false, err
		}

		if solved {
			return true, nil
		} else {
			// Reset for Next Itteration
			b.ResetValue(row, col)
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

func clearScreen() {
	switch runtime.GOOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default: // linux, mac, etc.
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}
