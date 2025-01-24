package board

import (
	"errors"
	"fmt"
)

// TODO: Implement feature for Possibility Space
// TODO: Implement feature for Fixed/Locked Cells for board reset (For Gamification)
type Cell struct {
	Value         int
	Possibilities []int
	Fixed         bool
}

type Board struct {
	cells  [9][9]Cell
	solved bool
}

type BoardInterface interface {
	GetValue(row, col int) (int, error)
	SetValue(row, col, value int) error
	IsValidMove(row, col, value int) (bool, error)
	IsSolved() (bool, error)
	Print()
}

func NewBoard() *Board {
	return &Board{
		cells:  [9][9]Cell{},
		solved: false,
	}
}

func (b *Board) GetValue(row, col int) (int, error) {
	if err := validatePosition(row, col); err != nil {
		return 0, err
	}
	cell := b.cells[row][col]
	return cell.Value, nil
}

func (b *Board) SetValue(row, col, value int) error {

	if err := validatePosition(row, col); err != nil {
		return err
	}
	if err := validateValue(value); err != nil {
		return err
	}

	b.cells[row][col].Value = value
	b.cells[row][col].Fixed = true
	// TODO: Update posibilities of Row, Col & Box
	return nil
}
func (b *Board) ResetValue(row, col int) error {

	if err := validatePosition(row, col); err != nil {
		return err
	}

	b.cells[row][col].Value = 0
	b.cells[row][col].Fixed = false
	return nil
}

// Check Unique Val in Row
// Check Unique Val in Col
// Check Unique Val in Box
func (b *Board) IsValidMove(row, col, value int) (bool, error) {
	if err := validatePosition(row, col); err != nil {
		return false, err
	}
	if err := validateValue(value); err != nil {
		return false, err
	}

	if !b.validateRow(row, value) {
		return false, nil
	}
	if !b.validateCol(col, value) {
		return false, nil
	}
	if !b.validateBox(row, col, value) {
		return false, nil
	}

	// TODO: Implement Check
	return true, nil
}

func (b *Board) IsSolved() (bool, error) {
	// TODO: Implement Solved Check
	return false, errors.New("STUB: Function not implemented")
}

func (b *Board) Print() {
	fmt.Println("===============")
	for i, row := range b.cells {
		for j, cell := range row {
			if (j)%3 == 0 {
				fmt.Print("|")
			}
			fmt.Print(cell.Value)

		}
		// +1 Offset to handle array starting at 0
		if (i+1)%3 == 0 {
			fmt.Println()
			fmt.Println("-------------")
		} else {
			fmt.Println()
		}
	}
	fmt.Println("===============")
}

// ==============================================
// 					Update Possibilities
// ==============================================

func (b *Board) UpdateAllBoardPossibilities() error {
	for i, row := range b.cells {
		for j := range row {
			b.updatePossibilities(i, j)
		}
	}
	return nil
}

func (b *Board) updatePossibilities(row, col int) error {
	// Escape case for already set values
	if b.cells[row][col].Fixed {
		return nil
	}

	//Get all Values from Row
	values := []int{}
	for i := 0; i < 9; i++ {
		val, err := b.GetValue(row, i)
		if err != nil {
			return err
		}
		values = append(values, val)
	}

	//Get all values from Col
	for i := 0; i < 9; i++ {
		val, err := b.GetValue(i, col)
		if err != nil {
			return err
		}
		values = append(values, val)
	}

	//Get all values from Box
	box_row := (row / 3) * 3
	box_col := (col / 3) * 3
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			rx := box_row + r
			cx := box_col + c
			val, err := b.GetValue(rx, cx)
			if err != nil {
				return err
			}
			values = append(values, val)
		}
	}

	// TODO: Compare Values to 1-9 Array and find whats missing.
	// 	- Whatever is missing = possibilities
	//	- Save Possibilities to Cell
	possibilities := findMissingNumbers(values)
	b.cells[row][col].Possibilities = possibilities

	return nil
}

func findMissingNumbers(numbers []int) []int {
	// Create a map to track seen numbers
	seen := make(map[int]bool)

	// Mark all numbers we've seen
	for _, num := range numbers {
		if num != 0 { // Ignore empty cells (0s)
			seen[num] = true
		}
	}

	// Collect missing numbers
	missing := []int{}
	for i := 1; i <= 9; i++ {
		if !seen[i] {
			missing = append(missing, i)
		}
	}

	return missing
}

// ==============================================
// 					Validators
// ==============================================

func validatePosition(row, col int) error {
	if row < 0 || row >= 9 || col < 0 || col >= 9 {
		return newPositionError(row, col)
	}
	return nil
}

func validateValue(value int) error {
	if value < 0 || value > 9 {
		return newValueError(value)
	}
	return nil
}

func (b *Board) validateRow(row, value int) bool {
	for col := 0; col < 9; col++ {
		if b.cells[row][col].Value == value {
			return false
		}
	}
	return true
}

func (b *Board) validateCol(col, value int) bool {
	for row := 0; row < 9; row++ {
		if b.cells[row][col].Value == value {
			return false
		}
	}
	return true
}

func (b *Board) validateBox(row, col, value int) bool {
	box_row := (row / 3) * 3
	box_col := (col / 3) * 3

	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			rx := box_row + r
			cx := box_col + c

			// Skip the current position we're validating
			if rx == row && cx == col {
				continue
			}

			if b.cells[rx][cx].Value == value {
				return false
			}
		}
	}
	return true
}
