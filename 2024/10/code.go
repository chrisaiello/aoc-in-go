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

type Coord struct {
	X int
	Y int
}

var steps = map[string]Coord{
	"up":    {-1, 0},
	"down":  {1, 0},
	"left":  {0, -1},
	"right": {0, 1},
}

var endings = make(map[Coord]int)

var score int = 0

func takestep(trlmap [][]int, pos Coord) int {
	X := pos.X
	Y := pos.Y
	curr_height := trlmap[X][Y]
	//fmt.Println("POS:", pos)
	//fmt.Println("curr_height:", curr_height)
	if curr_height == 9 {
		fmt.Println("DONE! Reached 9 at", pos)
		//endings[pos]++
		score++
		return 1
	}

	for k, s := range steps {
		if X+s.X < 0 || Y+s.Y < 0 {
			//fmt.Println("Outside of map-")
			continue
		}
		if X+s.X >= len(trlmap) || Y+s.Y >= len(trlmap) {
			//fmt.Println("Outside of map+")
			continue
		}

		next_height := trlmap[X+s.X][Y+s.Y]
		if next_height == curr_height+1 {
			fmt.Println("Moving:", k, "from", pos)
			//fmt.Println("next_height:", next_height)
			takestep(trlmap, Coord{X + s.X, Y + s.Y})
		}
	}

	return 0
}

func run(part2 bool, input string) any {

	trlmap := make([][]int, 0)
	for _, row := range strings.Split(input, "\n") {
		intRow := make([]int, 0, len(row))
		for _, ch := range row {
			digit, _ := strconv.Atoi(string(ch))
			intRow = append(intRow, digit)
		}
		trlmap = append(trlmap, intRow)
	}

	for i, row := range trlmap {
		for j, el := range row {
			if el == 0 {
				fmt.Println("TRAILHEAD:", i, j)
				takestep(trlmap, Coord{i, j})
				//fmt.Println(endings)
				//score += len(endings)
				//endings = make(map[Coord]int)
			}
		}
	}

	return score
}
