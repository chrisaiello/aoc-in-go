package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
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
					for z := 1; z < 3; z++ {
						//fmt.Println("x:", x, "y:", y, "z:", z)
						//fmt.Println("Looking for:", MAS[z], "in", dir)

						if x_new < 0 || y_new < 0 {
							//fmt.Println("BRK < 0")
							//count++
							break
						}

						if x_new > len(grid)-1 || y_new > len(grid)-1 {
							//fmt.Println("BRK out of grid")
							//count++
							break
						}

						if grid[x_new][y_new] != MAS[z] {
							//fmt.Println("BRK wrong letter")
							//count++
							break
						}
						if MAS[z] == "A" && grid[x_new][y_new] == "A" {
							key = fmt.Sprintf("%d,%d", x_new, y_new)
						}
						// If we found the "A" in "MAS",
						// record its coordinates
						// Later, check for multiple ocurrences of the
						// same "A" coordinate

						if z == 2 {
							fmt.Println("FOUND MAS")
							m[key] += 1
							count++
						}

						x_new += dir[0]
						y_new += dir[1]

					}
				}

			}
		}
	}

	//fmt.Printf("%T", grid)
	//fmt.Println(grid[9][19])
	fmt.Println(m)
	x_count := 0
	for key, value := range m {
		if value == 2 {
			fmt.Println(key, value)
			x_count++
		}
	}
	return x_count
}

// func run(part2 bool, input string) any {
// 	XMAS := []string{"X", "M", "A", "S"}
// 	lines := strings.Split(input, "\n")
// 	grid := make([][]string, len(lines))
// 	for i := range grid {
// 		grid[i] = strings.Split(lines[i], "")
// 	}

// 	dirs := [][]int{
// 		// {x,y} search directions
// 		{-1, -1},
// 		{-1, 0},
// 		{-1, 1},
// 		{0, -1},
// 		{0, 1},
// 		{1, -1},
// 		{1, 0},
// 		{1, 1},
// 	}
// 	count := 0
// 	for x := range grid {
// 		for y := range grid {
// 			//count++
// 			if grid[x][y] == XMAS[0] {
// 				//count++
// 				for _, dir := range dirs {
// 					x_new := x + dir[0]
// 					y_new := y + dir[1]
// 					for z := 1; z < 4; z++ {
// 						fmt.Println("x:", x, "y:", y, "z:", z)
// 						fmt.Println("Looking for:", XMAS[z], "in", dir)

// 						if x_new < 0 || y_new < 0 {
// 							fmt.Println("BRK < 0")
// 							//count++
// 							break
// 						}

// 						if x_new > len(grid)-1 || y_new > len(grid)-1 {
// 							fmt.Println("BRK out of grid")
// 							//count++
// 							break
// 						}

// 						if grid[x_new][y_new] != XMAS[z] {
// 							fmt.Println("BRK wrong letter")

// 							//count++
// 							break
// 						}
// 						if z == 3 {
// 							fmt.Println("FOUND XMAS")
// 							count++
// 						}

// 						x_new += dir[0]
// 						y_new += dir[1]

// 					}
// 				}

// 			}
// 		}
// 	}

// 	//fmt.Printf("%T", grid)
// 	//fmt.Println(grid[9][19])
// 	return count
// }
