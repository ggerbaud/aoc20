package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "2"

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
		p, pass := getPasswordAndPolicy(line)
		if verifyPasswordAndPolicy(pass, p) {
			total++
		}
	}
	return total
}

func part2(lines []string) int {
	total := 0
	for _, line := range lines {
		p, pass := getPasswordAndPolicy(line)
		if verifyPasswordAndPolicy2(pass, p) {
			total++
		}
	}
	return total
}

func getPasswordAndPolicy(line string) (policy, string) {
	p, pass, _ := strings.Cut(line, ": ")
	minmax, letter, _ := strings.Cut(p, " ")
	minStr, maxStr, _ := strings.Cut(minmax, "-")
	return policy{min: utils.ParseInt(minStr), max: utils.ParseInt(maxStr), letter: rune(letter[0])}, pass
}
func verifyPasswordAndPolicy(password string, p policy) bool {
	count := 0
	for _, c := range password {
		if c == p.letter {
			count++
		}
	}
	return count >= p.min && count <= p.max
}

func verifyPasswordAndPolicy2(password string, p policy) bool {
	return (rune(password[p.min-1]) == p.letter) != (rune(password[p.max-1]) == p.letter)
}

type policy struct {
	min, max int
	letter   rune
}
