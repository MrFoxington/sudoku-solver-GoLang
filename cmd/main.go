package main

import (
	"log"
	Board "sudoku-solver/internal/board"
)

func main() {

	err := run()
	if err != nil {
		log.Println("IT BROKE")
	}
}

// Add Code here to run Solver
func run() error {

	b := Board.NewBoard()

	b.Print()
	return nil
}
