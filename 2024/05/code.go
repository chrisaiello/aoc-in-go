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

func run(part2 bool, input string) any {
	parts := strings.Split(input, "\n\n")
	rules_s := strings.Split(parts[0], "\n")
	books_s := strings.Split(parts[1], "\n")

	rules := make([][]int, len(rules_s))
	for i, rule := range rules_s {
		s1 := strings.Split(rule, "|")
		x1, _ := strconv.Atoi(s1[0])
		x2, _ := strconv.Atoi(s1[1])
		rules[i] = []int{x1, x2}
	}

	books := make([][]int, len(books_s))
	for i, book := range books_s {
		s1 := strings.Split(book, ",")
		s2 := make([]int, len(s1))
		for j, s := range s1 {
			x, _ := strconv.Atoi(s)
			s2[j] = x
		}
		books[i] = s2
	}
	fmt.Println("Rules:\n", rules, "\n...")
	fmt.Println("Books:\n", books)
	return 1
}
