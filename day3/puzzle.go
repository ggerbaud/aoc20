package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "3"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	return countTreesForSlope(lines, 3, 1)
}

func part2(lines []string) int {
	a := countTreesForSlope(lines, 1, 1)
	b := countTreesForSlope(lines, 3, 1)
	c := countTreesForSlope(lines, 5, 1)
	d := countTreesForSlope(lines, 7, 1)
	e := countTreesForSlope(lines, 1, 2)
	return a * b * c * d * e
}

func countTreesForSlope(lines []string, right int, down int) int {
	total := 0
	x := 0
	for y := 0; y < len(lines); y += down {
		line := lines[y]
		if line[x] == '#' {
			total++
		}
		x = (x + right) % len(line)
	}
	return total
}
