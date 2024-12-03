package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func run(part2 bool, input string) any {
	// mul(11,8)
	re1 := regexp.MustCompile(`mul\((\d+,\d+)\)`)
	re2 := regexp.MustCompile(`(?s)don't\(\).*?do\(\)`)
	sum := 0

	// trim := strings.TrimSpace(input)
	// lines := strings.Split(trim, "\n")

	//for _, line := range lines {
	cleaned_line := re2.ReplaceAllString(input, "")
	matches := re1.FindAllStringSubmatch(cleaned_line, -1)
	//matches2 := re2.FindAllString(line, -1)
	//fmt.Println(line)
	fmt.Println(cleaned_line)
	//fmt.Println(matches2)
	for _, match := range matches {
		nums := strings.Split(match[1], ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		sum += x * y

	}
	//}
	return sum
}

// Part 1
// func run(part2 bool, input string) any {
// 	if part2 {
// 		return "not implemented"
// 	}
// 	// mul(11,8)
// 	re1 := regexp.MustCompile(`mul\((\d+,\d+)\)`)
// 	sum := 0

// 	trim := strings.TrimSpace(input)
// 	lines := strings.Split(trim, "\n")

// 	for _, line := range lines {
// 		matches := re1.FindAllStringSubmatch(line, -1)
// 		fmt.Println(matches)
// 		for _, match := range matches {
// 			nums := strings.Split(match[1], ",")
// 			x, _ := strconv.Atoi(nums[0])
// 			y, _ := strconv.Atoi(nums[1])
// 			sum += x * y

// 		}
// 	}
// 	return sum
// }
