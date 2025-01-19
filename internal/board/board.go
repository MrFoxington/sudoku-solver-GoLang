package Board

import (
	"errors"
	"fmt"
)

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
	GetValue(row int, col int) (int, error)
	SetValue(row int, col int, value int) error
	IsValidMove(row int, col int, value int) bool
	IsSolved() bool
	Print()
}

func NewBoard() *Board {
	return &Board{
		cells:  [9][9]Cell{},
		solved: false,
	}
}

func (b *Board) GetValue(row, col int) (int, error) {
	cell := b.cells[row][col]
	return cell.Value, nil
}

func (b *Board) SetValue(row, col, value int) error {
	b.cells[row][col].Value = value
	return nil
}

func (b *Board) IsValidMove(row, col, value int) (bool, error) {
	// TODO: Implement Check
	return false, errors.New("STUB: Function not implemented")
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
