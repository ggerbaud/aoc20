package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "1"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	exp := getExpenses(lines)
	for i := 0; i < len(exp)-1; i++ {
		for j := i + 1; j < len(exp); j++ {
			if exp[i]+exp[j] == 2020 {
				return exp[i] * exp[j]
			}
		}
	}
	return 0
}

func part2(lines []string) int {
	exp := getExpenses(lines)
	for i := 0; i < len(exp)-2; i++ {
		for j := i + 1; j < len(exp)-1; j++ {
			for k := j + 1; k < len(exp); k++ {
				if exp[i]+exp[j]+exp[k] == 2020 {
					return exp[i] * exp[j] * exp[k]
				}
			}
		}
	}
	return 0
}

func getExpenses(lines []string) []int {
	expenses := make([]int, len(lines))
	for i, line := range lines {
		expenses[i] = utils.ParseInt(line)
	}
	return expenses
}
