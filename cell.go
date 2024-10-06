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

func (g *Grid) fillRandom() {

	for i := 0; i < len(g.cells); i++ {
		for j := 0; j < len(g.cells[i]); j++ {
			randBool := rand.Intn(2) == 0

			fmt.Printf("(%d, %d) - %t\n", i, j, randBool)
			g.cells[i][j] = *NewCell(i, j, randBool)
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
	row        int
	col        int
	alive      bool
	nextState  bool
	grid       [][]Cell
	generation chan int
	mutex      *sync.Mutex
	barrier    *sync.WaitGroup
}

func NewCell(row, col int, alive bool) *Cell {
	return &Cell{row: row, col: col, alive: alive, mutex: &sync.Mutex{}, barrier: &sync.WaitGroup{}, generation: make(chan int)}
}

func (c *Cell) Run() {
	for gen := range c.generation {
		c.ComputeNextState()
		c.barrier.Done()
		c.barrier.Wait()
		c.UpdateState()
		c.barrier.Done()
		if gen == GENERATIONS-1 {
			close(c.generation)
		}
	}
}

func (c *Cell) ComputeNextState() {
	aliveNeighbours := c.CountAliveNeighbours()

	if c.alive {
		c.nextState = aliveNeighbours == 2 || aliveNeighbours == 3
	} else {
		c.nextState = aliveNeighbours == 3
	}
}

func (c *Cell) UpdateState() {
	c.mutex.Lock()
	c.alive = c.nextState
	c.mutex.Unlock()
}

func (c *Cell) isAlive() bool {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	return c.alive
}

func (c *Cell) CountAliveNeighbours() int {
	count := 0
	for i := c.row - 1; i <= c.row+1; i++ {
		for j := c.row - 1; j <= c.row+1; j++ {
			if i == c.row && j == c.row {
				continue
			}

			if i == -1 || j == -1 {
				continue
			}

			if c.grid[i][j].isAlive() {
				count++
			}
		}
	}
	return count
}
