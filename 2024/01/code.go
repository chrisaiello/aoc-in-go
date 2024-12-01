package main

import (
	"github.com/jpillora/puzzler/harness/aoc"
	"sort"
	"strings"
	"strconv"
)

func main() {
	aoc.Harness(run)
}

// on code change, run will be executed 4 times:
// 1. with: false (part1), and example input
// 2. with: true (part2), and example input
// 3. with: false (part1), and user input
// 4. with: true (part2), and user input
// the return value of each run is printed to stdout
func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block
	if part2 {
		return "not implemented"
	}
	// split input into lines
	lines := strings.Split(strings.TrimSpace(input), "\n")
	
	// create slices for each column
	col1 := make([]int, 0, len(lines))
	col2 := make([]int, 0, len(lines))
	
	// parse each line
	for _, line := range lines {
		// split and trim spaces
		fields := strings.Fields(line)
		if len(fields) != 2 {
			continue
		}
		
		// parse numbers
		n1, err := strconv.Atoi(fields[0])
		if err != nil {
			continue
		}
		n2, err := strconv.Atoi(fields[1])
		if err != nil {
			continue
		}
		
		col1 = append(col1, n1)
		col2 = append(col2, n2)
	}
	
	// sort both slices
	sort.Ints(col1)
	sort.Ints(col2)
	
	return len(col1) // temporary return value
}
