package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run1)
}

var page_number_map map[int]int

func validate_book(book []int) (bool, int) {
	nbook := make([]int, len(book))
	copy(nbook, book)
	mapped_book := map_book(nbook)
	if ascending(mapped_book) {
		return true, book[len(book)/2]
	} else {
		return false, 0
	}
}

func map_book(book []int) []int {
	for i, page := range book {
		mapped_page := page_number_map[page]
		book[i] = mapped_page
	}
	return book
}

func contains(candidate int, rules [][]int) bool {
	for _, rule := range rules {
		if candidate == rule[1] {
			fmt.Println("Candidate", candidate, "was found in", rule)
			return false
		}
	}
	fmt.Println("This must be the lowest")
	return true
}

func prune(x int, rules [][]int) [][]int {
	rules = slices.DeleteFunc(rules, func(n []int) bool {
		return n[0] == x
	})
	return rules
}

func gen_page_number_map(rules [][]int) any {
	fmt.Println("starting")
	fmt.Println(len(rules))
	if len(rules) == 1 {
		fmt.Println("Reached last element")
		page_number_map[rules[0][0]] = len(page_number_map)
		page_number_map[rules[0][1]] = len(page_number_map)
		return true
	}

	for _, rule := range rules {

		if len(rules) == 0 || rule == nil {
			fmt.Println("Reached end?")
			return true
		}
		candidate := rule[0]
		//fmt.Println("Candidate:", candidate)
		is_lowest := contains(candidate, rules)
		if is_lowest {
			page_number_map[candidate] = len(page_number_map)
			//fmt.Println("Rules:", rules, "Removing:", candidate)
			rules = prune(candidate, rules)
			//fmt.Println("New rules:", rules)
			gen_page_number_map(rules)
		}

	}

	return true
}

func ascending(s []int) bool {
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			return true
		}
		if s[i] >= s[i+1] {
			return false
		}
	}
	return false
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
	page_number_map = make(map[int]int)
	fmt.Println("Taking", rules[0][0], "as starting point")
	page_number_map[rules[0][0]] = len(page_number_map)
	rules = prune(rules[0][0], rules)
	gen_page_number_map(rules)
	//fmt.Println(ascending([]int{1, 5, 6, 5}))
	fmt.Println(page_number_map)
	sum := 0
	for _, book := range books {
		_, mid := validate_book(book)
		sum += mid
		//fmt.Println(valid, mid)
	}
	//fmt.Println("Rules:\n", rules, "\n...")
	//fmt.Println("Books:\n", books)
	return sum
}
func rules_contain1(rules [][]int, test_rule []int) bool {
	for _, rule := range rules {
		if rule[0] == test_rule[0] && rule[1] == test_rule[1] {
			return true
		}
	}
	return false
}

// Part 1
func validate_book1(book []int, rules [][]int) (bool, int) {
	for i, outer_page := range book {
		for _, check_page := range book[i+1:] {
			if rules_contain1(rules, []int{check_page, outer_page}) {
				return false, 0
			}
		}
	}
	return true, book[len(book)/2]
}

func get_sorted_middle2(book []int, rules [][]int) int {
	len := len(book)
	matches := 0
	fmt.Println(book)
	for _, outer_page := range book {
		for _, check_page := range book {
			if rules_contain1(rules, []int{check_page, outer_page}) {
				matches++
			}
		}
		if matches == len/2 {
			return outer_page
		}
		matches = 0
	}
	return 0
}

// Part 2
func validate_book2(book []int, rules [][]int) (bool, int) {
	for i, outer_page := range book {
		for _, check_page := range book[i+1:] {
			if rules_contain1(rules, []int{check_page, outer_page}) {
				return false, get_sorted_middle2(book, rules)
			}
		}
	}
	return true, 0
}

func run1(part2 bool, input string) any {
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
	sum := 0
	for _, book := range books {
		valid, mid := validate_book2(book, rules)
		sum += mid
		fmt.Println(valid, mid)
	}

	return sum
}
