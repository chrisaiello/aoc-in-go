package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

var safe int = 0

type Row struct {
	row   []int
	retry bool
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}
func remove(slice []int, s int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice[:s]...)
	newSlice = append(newSlice, slice[s+1:]...)
	return newSlice
}

func increasing_or_decreasing(row Row) bool {
	increasing := true
	decreasing := true
	// fail1 := -1
	// fail2 := -1

	for i, val := range row.row {
		if i < len(row.row)-1 && row.row[i+1] >= val {
			//fmt.Println(row[i+1], ">=", val)

			decreasing = false
		}
	}
	for i, val := range row.row {
		if i < len(row.row)-1 && row.row[i+1] <= val {
			increasing = false
		}
	}
	//fmt.Println(row, "increasing: ", increasing, "decreasing: ", decreasing)
	return increasing || decreasing
}

func range_check(row Row) bool {
	for i, val := range row.row {
		if i == len(row.row)-1 {
			// if row.retry {
			// 	fmt.Println(row, "is safe")
			// }
			return true
		} else if abs(val-row.row[i+1]) > 3 {
			return false
		}
	}
	return false
}

func validate_row(row Row) bool {
	if range_check(row) && increasing_or_decreasing(row) {
		if row.retry {
			fmt.Println(row, "is safe")
		}
		if !row.retry {
			safe++
		}
		return true
	}
	if !row.retry {
		retry_success := false
		for i := range row.row {
			new_row := Row{remove(row.row, i), true}
			fmt.Println("Retrying ", row, "with: ", new_row)
			if validate_row(new_row) {
				retry_success = true
			}
		}
		if retry_success {
			safe++
		}
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
		row := Row{int_s, false}

		validate_row(row)

	}

	return safe // temporary return value
}
