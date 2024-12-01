package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
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

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block

	// split input into lines
	lines := strings.Split(strings.TrimSpace(input), "\n")

	// create slices for each column
	col1 := make([]int, 0, len(lines))
	col2 := make([]int, 0, len(lines))
	diffs := make([]int, 0, len(lines))

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

	// define a map of frequencies

	freq := make(map[int]int)

	if part2 {
		sum := 0
		// populate frequency map
		for _, v := range col2 {
			freq[v] += 1
		}

		for _, v := range col1 {
			sum += v * freq[v]
		}

		return sum
	}

	for i := range col1 {
		x := col1[i]
		y := col2[i]
		diff := abs(x - y)

		fmt.Printf("x: %d y: %d d: %d \n", x, y, diff)

		diffs = append(diffs, diff)
	}

	sum := 0

	for _, v := range diffs {
		sum += v
	}

	return sum // temporary return value
}
