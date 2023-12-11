package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "6"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	group := make(map[string]bool)
	for _, line := range lines {
		if line == "" {
			total += len(group)
			group = make(map[string]bool)
			continue
		}
		for _, c := range line {
			group[string(c)] = true
		}
	}
	total += len(group)
	return total
}

func part2(lines []string) int {
	total := 0
	size := 0
	group := make(map[string]int)
	for _, line := range lines {
		if line == "" {
			for _, v := range group {
				if v == size {
					total++
				}
			}
			group = make(map[string]int)
			size = 0
			continue
		}
		size++
		for _, c := range line {
			v, ok := group[string(c)]
			if ok {
				group[string(c)] = v + 1
			} else {
				group[string(c)] = 1
			}
		}
	}
	for _, v := range group {
		if v == size {
			total++
		}
	}
	return total
}
