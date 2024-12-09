package main

import (
	"strings"
	"sync"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Coord struct {
	X int
	Y int
}

const (
	STEP_LIMIT = 7500
)

var step = map[string]Coord{
	"up":    {-1, 0},
	"down":  {1, 0},
	"left":  {0, -1},
	"right": {0, 1},
}

func rotate(dir string) string {
	switch dir {
	case "up":
		return "right"
	case "right":
		return "down"
	case "down":
		return "left"
	case "left":
		return "up"
	}
	return "err"
}

func find_start(grid [][]string) Coord {
	for i, row := range grid {
		for j, char := range row {
			if char == "^" {
				grid[i][j] = "X"
				return Coord{i, j}
			}
		}
	}
	return Coord{-1, -1}
}

// GridState encapsulates the mutable state for each move
type GridState struct {
	grid       [][]string
	guardLoc   Coord
	curDir     string
	totalSteps int
	foundLoop  bool
}

// deep copy of the grid
func copyGrid(original [][]string) [][]string {
	copied := make([][]string, len(original))
	for i := range original {
		copied[i] = make([]string, len(original[i]))
		copy(copied[i], original[i])
	}
	return copied
}

func move(state *GridState) {
	if state.totalSteps >= STEP_LIMIT {
		state.foundLoop = true
		return
	}

	moveStep := step[state.curDir]
	nextCoord := Coord{
		X: state.guardLoc.X + moveStep.X,
		Y: state.guardLoc.Y + moveStep.Y,
	}

	// Check grid boundaries
	if nextCoord.X < 0 || nextCoord.X >= len(state.grid) ||
		nextCoord.Y < 0 || nextCoord.Y >= len(state.grid[0]) {
		state.grid[state.guardLoc.X][state.guardLoc.Y] = "X"
		state.totalSteps++
		return
	}

	moveSquare := state.grid[nextCoord.X][nextCoord.Y]

	if moveSquare == "#" {
		state.curDir = rotate(state.curDir)
		move(state)
		return
	}

	state.totalSteps++
	state.grid[state.guardLoc.X][state.guardLoc.Y] = "X"
	state.guardLoc = nextCoord
	move(state)
}

func run(part2 bool, input string) any {
	lines := strings.Split(input, "\n")
	originalGrid := make([][]string, len(lines))
	for i := range originalGrid {
		originalGrid[i] = strings.Split(lines[i], "")
	}

	// Use WaitGroup to coordinate goroutines
	var wg sync.WaitGroup
	// Use Mutex to safely increment found_loops
	var mu sync.Mutex
	foundLoops := 0

	// Iterate through grid and run concurrent moves
	for i, row := range originalGrid {
		for j := range row {
			if originalGrid[i][j] == "." {
				wg.Add(1)
				go func(x, y int) {
					defer wg.Done()

					// Create a deep copy of the grid for this goroutine
					gridCopy := copyGrid(originalGrid)
					gridCopy[x][y] = "#"

					// Initialize state for this move
					state := &GridState{
						grid:       gridCopy,
						guardLoc:   find_start(gridCopy),
						curDir:     "up",
						totalSteps: 0,
						foundLoop:  false,
					}

					// Perform the move
					move(state)

					// If a loop was found, increment the counter safely
					if state.foundLoop {
						mu.Lock()
						foundLoops++
						mu.Unlock()
					}
				}(i, j)
			}
		}
	}

	// Wait for all goroutines to complete
	wg.Wait()

	return foundLoops
}
