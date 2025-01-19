package main

import (
	"fmt"
	"log"
	"sudoku-solver/internal/board"
	"sudoku-solver/internal/solver"
)

func main() {

	err := run()
	if err != nil {
		log.Fatalln("IT BROKE")
	}
}

// Add Code here to run Solver
func run() error {

	// var b board.BoardInterface = board.NewBoard()
	b := board.NewBoard()
	err := board.LoadPredefinedBoard(b, board.BoardEasy)
	if err != nil {
		return err

	}
	b.Print()

	s := solver.NewBacktrackingSolver()
	solved, err := s.Solve(b)
	if err != nil {
		log.Fatal(err)
	}
	if solved {
		fmt.Println("SOLVED!")
		b.Print()
	}
	return nil
}
