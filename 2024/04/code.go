package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run1)
	aoc.Harness(run2)
}

// Part 2
func run2(part2 bool, input string) any {
	MAS := []string{"M", "A", "S"}
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	for i := range grid {
		grid[i] = strings.Split(lines[i], "")
	}

	dirs := [][]int{
		// {x,y} search directions
		// MAS can only be diagonal
		{-1, -1},
		//{-1, 0},
		{-1, 1},
		//{0, -1},
		//{0, 1},
		{1, -1},
		//{1, 0},
		{1, 1},
	}
	count := 0
	m := make(map[string]int)
	for x := range grid {
		for y := range grid {
			if grid[x][y] == MAS[0] {
				for _, dir := range dirs {
					x_new := x + dir[0]
					y_new := y + dir[1]
					var key string
					for z := 1; z < len(MAS); z++ {

						if x_new < 0 || y_new < 0 {
							// We're outside of the grid
							break
						}

						if x_new > len(grid)-1 || y_new > len(grid)-1 {
							// We're outside of the grid
							break
						}

						if grid[x_new][y_new] != MAS[z] {
							// We didn't find the right letter
							break
						}
						if MAS[z] == "A" && grid[x_new][y_new] == "A" {
							// We found the middle "A" in the right spot
							// Record its index so we can find overlapping MAS's
							key = fmt.Sprintf("%d,%d", x_new, y_new)
						}

						if z == len(MAS)-1 {
							// We reached the end of "MAS"
							fmt.Println("FOUND MAS")
							// Save the coordinate of the "A"
							m[key] += 1
							count++
							break
						}

						// Continue searching in the same direction
						x_new += dir[0]
						y_new += dir[1]

					}
				}

			}
		}
	}

	fmt.Println(m)
	x_count := 0
	for key, value := range m {
		// Count instances where the same "A" coordinate is shared between multiple words
		if value == 2 {
			fmt.Println(key, value)
			x_count++
		}
	}
	return x_count
}

// Part 1
func run1(part2 bool, input string) any {
	XMAS := []string{"X", "M", "A", "S"}
	lines := strings.Split(input, "\n")
	grid := make([][]string, len(lines))
	for i := range grid {
		grid[i] = strings.Split(lines[i], "")
	}

	dirs := [][]int{
		// {x,y} search directions
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	count := 0
	for x := range grid {
		for y := range grid {
			if grid[x][y] == XMAS[0] {
				// Our initial check, is the current letter "X"?
				for _, dir := range dirs {
					// Search in every direction
					x_new := x + dir[0]
					y_new := y + dir[1]
					for z := 1; z < len(XMAS); z++ {
						fmt.Println("x:", x, "y:", y, "z:", z)
						fmt.Println("Looking for:", XMAS[z], "in", dir)

						if x_new < 0 || y_new < 0 {
							// We're outside of the grid
							break
						}

						if x_new > len(grid)-1 || y_new > len(grid)-1 {
							// We're outside of the grid
							break
						}

						if grid[x_new][y_new] != XMAS[z] {
							// We didn't find the right letter
							break
						}
						if z == len(XMAS)-1 {
							// We reached the end of XMAS
							fmt.Println("FOUND XMAS")
							count++
							break
						}
						// Continue searching in the same direction
						x_new += dir[0]
						y_new += dir[1]

					}
				}

			}
		}
	}

	return count
}
