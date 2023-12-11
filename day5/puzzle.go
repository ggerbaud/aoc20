package main

import (
	"advent/utils"
	"fmt"
	"strconv"
)

const day = "5"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	for _, line := range lines {
		n := decode(line)
		if n > total {
			total = n
		}
	}
	return total
}

func part2(lines []string) int {
	seats := make([]bool, 1024)
	for _, line := range lines {
		n := decode(line)
		seats[n] = true
	}
	for i := 1; i < 1023; i++ {
		if !seats[i] && seats[i-1] && seats[i+1] {
			return i
		}
	}
	return 0
}

func decode(line string) int {
	row := decodeRow(line[:7])
	col := decodeCol(line[7:])
	return row*8 + col
}

func decodeRow(line string) int {
	return decodeBin(line, 'F', 'B')
}

func decodeCol(line string) int {
	return decodeBin(line, 'L', 'R')
}

func decodeBin(line string, zero, one byte) int {
	result := 0
	for _, c := range line {
		if c == rune(zero) {
			result <<= 1
		} else if c == rune(one) {
			result = (result << 1) + 1
		}
	}
	return result
}
