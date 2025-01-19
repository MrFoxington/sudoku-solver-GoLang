package board

import (
	"errors"
	"fmt"
	"math"
)

// TODO: Implement feature for Possibility Space
// TODO: Implement feature for Fixed/Locked Cells for board reset (For Gamification)
type Cell struct {
	Value int
	// Possibilities []int
	// Fixed         bool
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

	if err := b.validateRow(row, value); err != nil {
		return false, err
	}
	if err := b.validateCol(col, value); err != nil {
		return false, err
	}
	if err := b.validateBox(row, col, value); err != nil {
		return false, err
	}

	// TODO: Implement Check
	return true, nil
}

func (b *Board) IsSolved() (bool, error) {
	// TODO: Implement Solved Check
	return false, errors.New("STUB: Function not implemented")
}

func (b *Board) Print() {
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

	box_row := int(math.Ceil(float64((row + 1) / 3)))
	box_col := int(math.Ceil(float64((col + 1) / 3)))

	for r := range 3 {
		for c := range 3 {
			rx := ((box_row * 3) + r)
			cx := ((box_col * 3) + c)

			if b.cells[rx][cx].Value == value {
				return false
			}

		}
	}
	return true
}
