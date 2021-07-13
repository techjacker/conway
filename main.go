package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("init grid:")
	iterations := 3
	grid := createGrid(5, 5)
	println("")
	for i := 0; i < iterations; i++ {
		fmt.Printf("\niteration %d", i)
		grid = processGrid(grid)
		println("")
	}
}

func countLiveNeighbours(grid [][]bool, i, j int) int {
	live := 0
	h := len(grid) - 1
	w := len(grid[0]) - 1
	for r := i - 1; r <= i+1; r++ {
		if r < 0 || r > h {
			continue
		}
		for c := j - 1; c <= j+1; c++ {
			if c < 0 || c > w {
				continue
			}
			if grid[r][c] == true {
				live++
			}
		}
	}
	return live
}

func processCell(grid [][]bool, i, j int) bool {
	liveNeighbours := countLiveNeighbours(grid, i, j)
	if grid[i][j] == true {
		if liveNeighbours == 2 || liveNeighbours == 3 {
			return true
		}
	} else {
		if liveNeighbours == 3 {
			return true
		}
	}
	return false
}

func processGrid(grid [][]bool) [][]bool {
	next := [][]bool{}
	for i := 0; i < len(grid); i++ {
		row := []bool{}
		for j := 0; j < len(grid[i]); j++ {
			cell := processCell(grid, i, j)
			row = append(row, cell)
		}
		fmt.Printf("\noutput %v", row)
		next = append(next, row)
	}
	println("")
	return next
}

func createGrid(h, w int) [][]bool {
	grid := [][]bool{}
	for i := 0; i < h; i++ {
		row := []bool{}
		for i := 0; i < w; i++ {
			row = append(row, randomState())
		}
		fmt.Printf("\ninput %v", row)
		grid = append(grid, row)
	}
	return grid
}
func randomState() bool {
	i := rand.Intn(100)
	if i > 50 {
		return true
	}
	return false
}
