package main

import (
	"advent/utils"
	"fmt"
	"strconv"
	"strings"
)

const day = "4"

func main() {
	lines := utils.ReadFileLinesForDay(day, false)
	total := part1(lines)
	fmt.Println("Day#" + day + ".1 : " + strconv.Itoa(total))
	total = part2(lines)
	fmt.Println("Day#" + day + ".2 : " + strconv.Itoa(total))
}

func part1(lines []string) int {
	total := 0
	p := make(passport)
	for _, line := range lines {
		if line == "" {
			if p.hasAllFileds() {
				total++
			}
			p = make(passport)
		} else {
			for _, data := range strings.Split(line, " ") {
				k, v, _ := strings.Cut(data, ":")
				p[k] = v
			}
		}
	}
	if p.hasAllFileds() {
		total++
	}
	return total
}

func part2(lines []string) int {
	total := 0
	p := make(passport)
	for _, line := range lines {
		if line == "" {
			if p.isValid() {
				total++
			}
			p = make(passport)
		} else {
			for _, data := range strings.Split(line, " ") {
				k, v, _ := strings.Cut(data, ":")
				p[k] = v
			}
		}
	}
	if p.isValid() {
		total++
	}
	return total
}

func (p passport) isValid() bool {
	return p.hasAllFileds() && p.hasValidFields()
}

func (p passport) hasAllFileds() bool {
	_, okbyr := p["byr"]
	_, okiyr := p["iyr"]
	_, okeyr := p["eyr"]
	_, okhgt := p["hgt"]
	_, okhcl := p["hcl"]
	_, okecl := p["ecl"]
	_, okpid := p["pid"]
	return okbyr && okiyr && okeyr && okhgt && okhcl && okecl && okpid
}

func (p passport) hasValidFields() bool {
	byr := p["byr"]
	if len(byr) != 4 {
		return false
	}
	byrInt := utils.ParseInt(byr)
	if byrInt < 1920 || byrInt > 2002 {
		return false
	}
	iyr := p["iyr"]
	if len(iyr) != 4 {
		return false
	}
	iyrInt := utils.ParseInt(iyr)
	if iyrInt < 2010 || iyrInt > 2020 {
		return false
	}
	eyr := p["eyr"]
	if len(eyr) != 4 {
		return false
	}
	eyrInt := utils.ParseInt(eyr)
	if eyrInt < 2020 || eyrInt > 2030 {
		return false
	}
	hgt := p["hgt"]
	if strings.HasSuffix(hgt, "cm") {
		hgtInt := utils.ParseInt(strings.TrimSuffix(hgt, "cm"))
		if hgtInt < 150 || hgtInt > 193 {
			return false
		}
	} else if strings.HasSuffix(hgt, "in") {
		hgtInt := utils.ParseInt(strings.TrimSuffix(hgt, "in"))
		if hgtInt < 59 || hgtInt > 76 {
			return false
		}
	} else {
		return false
	}
	hcl := p["hcl"]
	if len(hcl) != 7 || hcl[0] != '#' {
		return false
	}
	for _, c := range hcl[1:] {
		if !((c >= '0' && c <= '9') || (c >= 'a' && c <= 'f')) {
			return false
		}
	}
	ecl := p["ecl"]
	if ecl != "amb" && ecl != "blu" && ecl != "brn" && ecl != "gry" && ecl != "grn" && ecl != "hzl" && ecl != "oth" {
		return false
	}
	pid := p["pid"]
	if len(pid) != 9 {
		return false
	}
	for _, c := range pid {
		if !(c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}

type (
	passport map[string]string
)
