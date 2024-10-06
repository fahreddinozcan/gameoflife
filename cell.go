package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Grid struct {
	cells [][]Cell
}

func newGrid(rows int, cols int) *Grid {
	grid := make([][]Cell, rows)
	for i := range grid {
		grid[i] = make([]Cell, cols)
	}
	return &Grid{grid}
}

func (g *Grid) fillRandom(mutex *sync.Mutex, barrier *sync.WaitGroup) {

	for i := 0; i < len(g.cells); i++ {
		for j := 0; j < len(g.cells[i]); j++ {
			randBool := rand.Intn(2) == 0

			fmt.Printf("(%d, %d) - %t\n", i, j, randBool)
			g.cells[i][j] = *newCell(i, j, randBool, mutex, barrier)
		}
	}
}

func (g *Grid) print() {
	for i := 0; i < len(g.cells); i++ {
		for j := 0; j < len(g.cells[i]); j++ {
			if g.cells[i][j].alive {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
		fmt.Println()
	}
}

type Cell struct {
	row       int
	col       int
	alive     bool
	nextState bool
	mutex     *sync.Mutex
	barrier   *sync.WaitGroup
}

func newCell(row, col int, alive bool, mutex *sync.Mutex, barrier *sync.WaitGroup) *Cell {
	return &Cell{row: row, col: col, alive: alive, mutex: mutex, barrier: barrier}
}
