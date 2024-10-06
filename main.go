package main

import (
	"sync"
)

var (
	ROWS = 5
	COLS = 5
)

func main() {

	grid := newGrid(ROWS, COLS)

	mutex := new(sync.Mutex)
	barrier := new(sync.WaitGroup)
	grid.fillRandom(mutex, barrier)

	grid.print()

}
