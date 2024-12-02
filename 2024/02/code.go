package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

var safe int = 0

func main() {
	aoc.Harness(run)
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func increasing_or_decreasing(row []int) bool {
	increasing := true
	decreasing := true

	for i, val := range row {
		if i < len(row)-1 && row[i+1] >= val {
			//fmt.Println(row[i+1], ">=", val)
			decreasing = false
		}
	}

	for i, val := range row {
		if i < len(row)-1 && row[i+1] <= val {
			increasing = false
		}
	}
	//fmt.Println(row, "increasing: ", increasing, "decreasing: ", decreasing)
	return increasing || decreasing
}

func range_check(row []int) bool {
	for i, val := range row {
		if i == len(row)-1 {
			fmt.Println(row, "is safe")
			return true
		} else if abs(val-row[i+1]) > 3 {
			fmt.Println("Removing bad value")
			validate_row(remove(row, i))
			return false
		}
	}
	return false
}

func validate_row(row []int) bool {
	if range_check(row) && increasing_or_decreasing(row) {
		safe++
		return true
	}
	return false
}

func run(part2 bool, input string) any {
	// when you're ready to do part 2, remove this "not implemented" block

	// split input into lines
	trim := strings.TrimSpace(input)
	lines := strings.Split(trim, "\n")

	for _, line := range lines {
		split := strings.Split(line, " ")
		int_s := make([]int, 0, len(split))

		for _, char := range split {
			x, _ := strconv.Atoi(char)
			int_s = append(int_s, x)
		}

		validate_row(int_s)

	}

	return safe // temporary return value
}
