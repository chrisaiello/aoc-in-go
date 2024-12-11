package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//run("125 17")
	run("17639 47 3858 0 470624 9467423 5 188")
}

type RockInfo struct {
	rock   int
	blinks int
}

var rockCache = make(map[RockInfo]int)

func even(i int) bool {
	return len(strconv.Itoa(i))%2 == 0
}
func split(i int) (int, int) {
	s := strconv.Itoa(i)
	ln := len(s)
	left, _ := strconv.Atoi(s[:ln/2])
	right, _ := strconv.Atoi(s[ln/2:])
	//fmt.Println("LEFT:", left, "RIGHT", right)
	return left, right
}

func blink(stones []int) []int {
	new_stones := []int{}

	for _, stone := range stones {
		if stone == 0 {
			new_stones = append(new_stones, 1)
			continue

		} else if even(stone) {
			left, right := split(stone)
			//stone_count++
			new_stones = append(new_stones, left, right)

		} else {
			new_stones = append(new_stones, stone*2024)
		}

		//fmt.Println(stone)
	}
	//fmt.Println(new_stones)
	return new_stones
}

var stone_count int = 0

func blink2(stone int, blinks int) int {
	if blinks >= 100 {
		//fmt.Println("Reached blink limit: ", stone)
		return 0
	}

	if stone == 0 {
		//fmt.Println("Setting 0 -> 1")
		new_stone := 1
		return blink3(new_stone, blinks+1)
	} else if even(stone) {
		//fmt.Println("Splitting:", stone)
		left, right := split(stone)
		//fmt.Println("stone_count+ 1")
		//stone_count++
		leftResult := blink3(left, blinks+1)
		rightResult := blink3(right, blinks+1)
		return 1 + leftResult + rightResult
	} else {
		//fmt.Println("Multiplying", stone)
		new_stone := 2024 * stone
		return blink3(new_stone, blinks+1)
	}
}

func blink3(stone int, blinks int) int {
	val, ok := rockCache[RockInfo{stone, blinks}]
	if ok {
		//fmt.Println("FOUND", stone, blinks, val, "in cache!")
		//fmt.Println("stone_count+", val-1)
		//stone_count += val
		return val
	} else {
		//fmt.Println("Calculating", stone, blinks, "for cache.")
		result := blink2(stone, blinks)
		//fmt.Println("stone_count++", result)
		//stone_count += result
		//fmt.Println("Adding", stone, blinks, result, "to cache.")
		rockCache[RockInfo{stone, blinks}] = result
		return result
	}
}

func run(input string) any {

	// ssl := strings.Split(input, " ")
	// isl := make([]int, len(ssl))
	// for i, el := range ssl {
	// 	isl[i], _ = strconv.Atoi(el)
	// }

	// stones := make([]int, len(strings.Fields(input)))
	// for i, el := range strings.Fields(input) {
	// 	stones[i], _ = strconv.Atoi(el)
	// }

	// for i := 0; i < 25; i++ {
	// 	//fmt.Println("Blinking:", i)
	// 	//fmt.Println(stones)
	// 	stones = blink(stones)
	// }

	stones2 := make([]int, len(strings.Fields(input)))
	for i, el := range strings.Fields(input) {
		stones2[i], _ = strconv.Atoi(el)
	}
	stone_count += len(stones2)
	fmt.Println("Initializing stone_count to", stone_count)

	sum := 0
	for _, stone := range stones2 {
		count := blink3(stone, 0)
		fmt.Println("Stone", stone, "results in", count, "new stones")
		sum += count + 1
	}

	//fmt.Println(stones)
	//fmt.Println(rockCache)
	fmt.Println("Stone Count:", sum)
	//fmt.Println(stones)
	return 0
}
