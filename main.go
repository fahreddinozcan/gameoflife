package main

import (
	"fmt"
	"sync"
)

const (
	ROWS        = 5
	COLS        = 5
	GENERATIONS = 5
)

func main() {

	grid := newGrid(ROWS, COLS)

	grid.fillRandom()

	grid.print()

	for i := range grid.cells {
		for j := range grid.cells[i] {
			go grid.cells[i][j].Run()
		}
	}
	fmt.Println("HERE!!!")

	for gen := 0; gen < GENERATIONS; gen++ {
		computeBarrier := &sync.WaitGroup{}
		updateBarrier := &sync.WaitGroup{}
		computeBarrier.Add(ROWS * COLS)
		updateBarrier.Add(ROWS * COLS)

		fmt.Printf("GENERATION %d\n", gen)
		for i := range grid.cells {
			for j := range grid.cells[i] {
				fmt.Printf("%d %d\n", i, j)
				grid.cells[i][j].computeBarrier = computeBarrier
				grid.cells[i][j].updateBarrier = updateBarrier
				fmt.Printf("BLOCK %d %d\n", i, j)
				grid.cells[i][j].generation <- gen
			}
		}
		fmt.Printf("GENERATION %d STEP-1\n", gen)

		computeBarrier.Wait()
		updateBarrier.Wait()

		grid.print()
		fmt.Printf("Generation of %d cells\n", len(grid.cells))

	}
}
