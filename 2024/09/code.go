package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/jpillora/puzzler/harness/aoc"
)

func main() {
	aoc.Harness(run)
}

func check_ordered1(layout []string) bool {
	dot_index := slices.Index(layout, ".")

	for i := dot_index; i < len(layout); i++ {
		if layout[i] != "." {
			//fmt.Println("Not sorted yet")
			return false
		}
	}
	fmt.Println("We're sorted!")
	return true
}

func move1(layout []string) ([]string, bool) {

	if check_ordered1(layout) {
		return layout, true
	}

	id_to_move := -1
	//fmt.Println(len(layout))
	for i := len(layout) - 1; i >= 0; i-- {
		//fmt.Println("Checking", layout[i], "at", i)
		if layout[i] != "." {
			//fmt.Println("Moving", layout[i], "at", i)
			id_to_move = i
			break
		}
	}

	for i, el := range layout {
		if el == "." {
			layout[i] = layout[id_to_move]
			layout[id_to_move] = "."
			break
		}
	}
	return layout, false
}

func score(layout []string) int {
	sum := 0
	for i, el := range layout {
		el_int := 0
		if el != "." {
			el_int, _ = strconv.Atoi(el)
		}
		sum += el_int * i
	}
	return sum
}

func has_free_space(layout []string, file_size int, candidate_index int) (bool, int) {
	candidate := -1
	size := 0

	for i, el := range layout {
		if i == candidate_index {
			return false, -1
		}
		if el == "." {
			candidate = i
			size++
			for j := i + 1; j < len(layout); j++ {
				if layout[j] == "." {
					//fmt.Println("size", size, "++ing")
					size++
				} else {
					break
				}
			}
			if size >= file_size {
				//fmt.Println(layout)
				fmt.Println("Layout has free space size", size, "at index", candidate)
				return true, candidate
			}
		}
		size = 0
	}
	return false, -1
}

func move2(layout []string, char string) []string {

	fmt.Println("Moving", char)
	index := slices.Index(layout, char)
	file_size := 0
	for i := index; i <= len(layout); i++ {
		if i == len(layout) || layout[i] != char {
			break
		}
		file_size++
	}
	fmt.Println(char, "Begins at index", index, "size", file_size)

	has_free_space, free_space_index := has_free_space(layout, file_size, index)

	//fmt.Println(has_free_space, free_space_index)

	if !has_free_space {
		fmt.Println("No free space")
		return layout
	}

	for i := 0; i < file_size; i++ {
		//fmt.Println("Setting layout index", i, "to", char)
		layout[i+free_space_index] = char
		layout[i+index] = "."
	}
	//fmt.Println(layout)

	return layout
}

func run(part2 bool, input string) any {
	in := make([]int, 0)
	in_sum := 0
	for _, char := range strings.Split(input, "") {
		num, _ := strconv.Atoi(char)
		in = append(in, num)
		in_sum += num
	}
	//fmt.Println(in_sum)

	layout := make([]string, 0)
	//fmt.Println(len(layout))

	j := 0
	for i, char := range in {
		write_char := strconv.Itoa(j)
		if i%2 == 1 {
			write_char = "."
		} else {
			j++
		}
		for i := 0; i < char; i++ {
			layout = append(layout, write_char)
		}
	}
	//fmt.Println(layout)

	// we now have a []string, layout, with the disk layout
	// i.e. 0..111....22222
	// Input always seems to be odd, which means the layout always starts and ends with
	// a file, not empty space

	// done := false
	// for i := 0; done == false; i++ {
	// 	layout, done = move1(layout)
	// }
	for i := len(in) / 2; i >= 0; i-- {
		i_s := strconv.Itoa(i)
		//fmt.Println(layout)
		layout = move2(layout, i_s)
		//fmt.Println(layout)
	}

	return score(layout)
}
