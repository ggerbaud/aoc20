package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "7"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	bags := makeBags(lines)
	result := make(map[string]bool)
	data := make([]*bag, 1)
	data[0] = bags["shiny gold"]
	for len(data) > 0 {
		b := data[0]
		data = data[1:]
		for _, cb := range b.canBeContainedBy {
			if !result[cb.color] {
				result[cb.color] = true
				data = append(data, cb)
			}
		}
	}
	return len(result)
}

func part2(lines []string) int {
	bags := makeBags(lines)
	return bags["shiny gold"].countSubBags()
}

func makeBags(lines []string) map[string]*bag {
	bags := make(map[string]*bag)
	for _, line := range lines {
		color, contains, _ := strings.Cut(line, " bags contain ")
		b := getOrMakeBag(bags, color)
		if contains == "no other bags." {
			continue
		}
		for _, c := range strings.Split(contains, ", ") {
			n, cb, _ := strings.Cut(c, " ")
			clr, _, _ := strings.Cut(cb, " bag")
			cbag := getOrMakeBag(bags, clr)
			cbag.canBeContainedBy = append(cbag.canBeContainedBy, b)
			b.mustContain[cbag] = utils.ParseInt(n)
		}
	}
	return bags
}

func getOrMakeBag(bags map[string]*bag, color string) *bag {
	b, ok := bags[color]
	if !ok {
		b = &bag{color: color, mustContain: make(map[*bag]int), canBeContainedBy: make([]*bag, 0)}
		bags[color] = b
	}
	return b
}

func (b *bag) countSubBags() int {
	result := 0
	for cb, n := range b.mustContain {
		result += n + n*cb.countSubBags()
	}
	return result
}

type (
	bag struct {
		color            string
		mustContain      map[*bag]int
		canBeContainedBy []*bag
	}
)
