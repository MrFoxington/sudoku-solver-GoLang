package board

import "fmt"

// Error Definitions
type ValidationError struct {
	Field   string
	Problem string
	Value   interface{}
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("%s: %s (got: %v)", e.Field, e.Problem, e.Value)
}

// Error Constructors
func newPositionError(row, col int) error {
	return &ValidationError{
		Field:   "Position",
		Problem: "Invalid board coordinates",
		Value:   fmt.Sprintf("row=%d, col=%d", row, col),
	}
}

func newValueError(value int) error {
	return &ValidationError{
		Field:   "Value",
		Problem: "Value must be between 0 and 9",
		Value:   value,
	}
}

func newInvalidMoveError_Row(row, col, value int) error {
	return &ValidationError{
		Field:   "Row",
		Problem: "Value exists in Row",
		Value:   fmt.Sprintf("row=%d, col=%d, value=%d", row, col, value),
	}
}

func newInvalidMoveError_Col(row, col, value int) error {
	return &ValidationError{
		Field:   "Col",
		Problem: "Value exists in Col",
		Value:   fmt.Sprintf("row=%d, col=%d, value=%d", row, col, value),
	}
}

func newInvalidMoveError_Box(row, col, value int) error {
	return &ValidationError{
		Field:   "Box",
		Problem: "Value exists in Box",
		Value:   fmt.Sprintf("row=%d, col=%d, value=%d", row, col, value),
	}
}
