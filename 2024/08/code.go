package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Coord struct {
	X int
	Y int
}

// GenerateAlphanumericCharacters creates a slice containing all alphanumeric characters
func GenerateAlphanumericCharacters() []rune {
	var characters []rune

	// Add lowercase letters (a-z)
	for c := 'a'; c <= 'z'; c++ {
		characters = append(characters, c)
	}

	// Add uppercase letters (A-Z)
	for c := 'A'; c <= 'Z'; c++ {
		characters = append(characters, c)
	}

	// Add digits (0-9)
	for c := '0'; c <= '9'; c++ {
		characters = append(characters, c)
	}

	return characters
}

func get_antennas(char rune, grid [][]rune) []Coord {
	var coords []Coord
	for i, row := range grid {
		for j, val := range row {
			if val == char {
				coords = append(coords, Coord{i, j})
			}
		}
	}
	if len(coords) > 0 {
		fmt.Println("Char:", string(char), "Coords:", coords, len(coords))
	}
	return coords
}

func valid_node(coord Coord, antennas []Coord, grid [][]rune) bool {
	length := len(grid)

	if coord.X < 0 || coord.Y < 0 {
		//fmt.Println("x Outside grid- at", coord)
		return false
	}

	if coord.X >= length || coord.Y >= length {
		//fmt.Println("x Outside grid+ at", coord)
		return false
	}

	// if slices.Contains(antennas, coord) {
	// 	fmt.Println("x Node blocked by antenna at", coord)
	// 	return false
	// }

	fmt.Println("* Node is valid at", coord)
	return true
}

func count_nodes(antennas []Coord, grid [][]rune, countm map[Coord]int) {
	for i, antenna1 := range antennas {
		for _, antenna2 := range antennas[:i] {
			fmt.Println(antenna1, antenna2)
			dX := antenna1.X - antenna2.X
			dY := antenna1.Y - antenna2.Y

			node_candidates := make([]Coord, 0)
			search_range := int(float64(len(grid)) * 1.5)

			for n := 0; n <= search_range; n++ {
				node_candidates = append(node_candidates, Coord{antenna1.X - dX*n, antenna1.Y - dY*n})
				node_candidates = append(node_candidates, Coord{antenna2.X - dX*n, antenna2.Y - dY*n})
				node_candidates = append(node_candidates, Coord{antenna2.X + dX*n, antenna2.Y + dY*n})
				node_candidates = append(node_candidates, Coord{antenna1.X + dX*n, antenna1.Y + dY*n})
			}

			for _, candidate := range node_candidates {
				if valid_node(candidate, antennas, grid) {
					countm[candidate]++
				}
			}

		}
	}
}

func run(part2 bool, input string) any {
	countm := make(map[Coord]int)
	chars := GenerateAlphanumericCharacters()
	grid := make([][]rune, 0)
	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, []rune(row))
	}

	for _, char := range chars {
		antennas := get_antennas(char, grid)
		if len(antennas) > 0 {
			count_nodes(antennas, grid, countm)
		}
	}

	fmt.Println(countm)
	return len(countm)
}

/* Part 1 Archive

package main

import (
	"fmt"
	"slices"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

type Coord struct {
	X int
	Y int
}

// GenerateAlphanumericCharacters creates a slice containing all alphanumeric characters
func GenerateAlphanumericCharacters() []rune {
	var characters []rune

	// Add lowercase letters (a-z)
	for c := 'a'; c <= 'z'; c++ {
		characters = append(characters, c)
	}

	// Add uppercase letters (A-Z)
	for c := 'A'; c <= 'Z'; c++ {
		characters = append(characters, c)
	}

	// Add digits (0-9)
	for c := '0'; c <= '9'; c++ {
		characters = append(characters, c)
	}

	return characters
}

func get_antennas(char rune, grid [][]rune) []Coord {
	var coords []Coord
	for i, row := range grid {
		for j, val := range row {
			if val == char {
				coords = append(coords, Coord{i, j})
			}
		}
	}
	if len(coords) > 0 {
		fmt.Println("Char:", string(char), "Coords:", coords, len(coords))
	}
	return coords
}

func valid_node(coord Coord, antennas []Coord, grid [][]rune) bool {
	length := len(grid)

	if coord.X < 0 || coord.Y < 0 {
		fmt.Println("x Outside grid- at", coord)
		return false
	}

	if coord.X >= length || coord.Y >= length {
		fmt.Println("x Outside grid+ at", coord)
		return false
	}

	if slices.Contains(antennas, coord) {
		fmt.Println("x Node blocked by antenna at", coord)
		return false
	}

	fmt.Println("* Node is valid at", coord)
	return true
}

func count_nodes(antennas []Coord, grid [][]rune, countm map[Coord]int) {
	for i, antenna1 := range antennas {
		for _, antenna2 := range antennas[:i] {
			fmt.Println(antenna1, antenna2)
			dX := antenna1.X - antenna2.X
			dY := antenna1.Y - antenna2.Y

			node_candidate1 := Coord{antenna1.X + dX, antenna1.Y + dY}
			node_candidate2 := Coord{antenna1.X - dX, antenna1.Y - dY}
			node_candidate3 := Coord{antenna2.X + dX, antenna2.Y + dY}
			node_candidate4 := Coord{antenna2.X - dX, antenna2.Y - dY}

			if valid_node(node_candidate1, antennas, grid) {
				countm[node_candidate1]++
			}
			if valid_node(node_candidate2, antennas, grid) {
				countm[node_candidate2]++
			}
			if valid_node(node_candidate3, antennas, grid) {
				countm[node_candidate3]++
			}
			if valid_node(node_candidate4, antennas, grid) {
				countm[node_candidate4]++
			}

		}
	}
}

func run(part2 bool, input string) any {
	countm := make(map[Coord]int)
	chars := GenerateAlphanumericCharacters()
	grid := make([][]rune, 0)
	for _, row := range strings.Split(input, "\n") {
		grid = append(grid, []rune(row))
	}

	for _, char := range chars {
		antennas := get_antennas(char, grid)
		if len(antennas) > 0 {
			count_nodes(antennas, grid, countm)
		}
	}

	//fmt.Println(string(grid[1][8]))
	return len(countm)
}
*/
