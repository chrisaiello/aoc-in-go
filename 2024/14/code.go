package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jpillora/puzzler/harness/aoc"
)

const (
	W int = 101 //11
	H int = 103 //7
)

func main() {
	aoc.Harness(run)
}

type Coord struct {
	x int
	y int
}

type Guard struct {
	p Coord
	v Coord
}

var guards []Guard

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return x * -1
}

func mod(a, b int) int {
	return (a%b + b) % b
}

func score(guards []Guard) int {
	q1, q2, q3, q4 := 0, 0, 0, 0
	for _, guard := range guards {
		if guard.p.x < W/2 && guard.p.y < H/2 {
			// Top left
			q1++
		} else if guard.p.x > W/2 && guard.p.y < H/2 {
			// Top right
			q2++
		} else if guard.p.x < W/2 && guard.p.y > H/2 {
			// Botton left
			q3++
		} else if guard.p.x > W/2 && guard.p.y > H/2 {
			// Bottom right
			q4++
		}
	}
	fmt.Println(q1, q2, q3, q4)
	return q1 * q2 * q3 * q4
}

func move(guards []Guard) []Guard {
	var guards_moved []Guard
	for _, guard := range guards {
		t := 1

		x := mod(guard.p.x+guard.v.x*t, W)
		y := mod(guard.p.y+guard.v.y*t, H)

		guard_moved := Guard{Coord{x, y}, Coord{guard.v.x, guard.v.y}}
		guards_moved = append(guards_moved, guard_moved)
	}
	return guards_moved
}

func updateGrid(guards []Guard, i int) {
	var grid [H][W]string
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j] = "."
		}
	}

	for _, guard := range guards {
		grid[guard.p.y][guard.p.x] = "X"
	}

	if i%H == 39 {
		printGrid(grid)
		fmt.Println("SECONDS", i+1)
		time.Sleep(250 * time.Millisecond)
	}
}

func printGrid(grid [H][W]string) {
	// Determine a reasonable subset to print (e.g., first 20 rows and columns)
	rowsToPrint := 103
	colsToPrint := 101

	// Create a string builder for efficient string concatenation
	var sb strings.Builder

	// Print column numbers for reference
	//sb.WriteString("   ") // Indent for row numbers
	// for j := 0; j < colsToPrint; j++ {
	// 	sb.WriteString(fmt.Sprintf("%d", j))
	// }
	sb.WriteString("\n")

	// Print the grid
	for i := 0; i < rowsToPrint; i++ {
		// Print row number
		//sb.WriteString(fmt.Sprintf("%d ", i))

		// Print grid cells
		for j := 0; j < colsToPrint; j++ {
			sb.WriteString(fmt.Sprintf("%s", grid[i][j]))
		}
		sb.WriteString("\n")
	}

	// Print the full grid to the terminal
	fmt.Print(sb.String())
}

func run(part2 bool, input string) any {
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Fields(line)
		if len(parts) == 2 {
			p1, _ := strconv.Atoi(strings.Split(strings.TrimPrefix(parts[0], "p="), ",")[0])
			p2, _ := strconv.Atoi(strings.Split(strings.TrimPrefix(parts[0], "p="), ",")[1])
			v1, _ := strconv.Atoi(strings.Split(strings.TrimPrefix(parts[1], "v="), ",")[0])
			v2, _ := strconv.Atoi(strings.Split(strings.TrimPrefix(parts[1], "v="), ",")[1])

			guard := Guard{Coord{p1, p2}, Coord{v1, v2}}
			guards = append(guards, guard)
		}
	}

	for i := 0; i < 10000; i += 1 {
		guards = move(guards)
		updateGrid(guards, i)
		//fmt.Println("SECONDS:", i+1)
	}

	return 0
}
