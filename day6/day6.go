package day6

import (
	"strconv"
)

// Day6a solves the puzzle and returns a string
func Day6a(a []string) string {
	ltrs := make(map[rune]bool)
	total := 0
	for _, line := range a {
		if line == "" {
			total += calcTotal(ltrs)
			ltrs = make(map[rune]bool)
		} else {
			for _, i := range line {
				ltrs[i] = true
			}
		}
	}
	// Catch the last input
	total += calcTotal(ltrs)

	return strconv.Itoa(total)
}

// MaxUint is the max unsigned int
const MaxUint = ^uint(0)

// MaxInt is the max signed int
const MaxInt = int(MaxUint >> 1)

// Day6b solves the puzzle and returns a string
func Day6b(a []string) string {
	ltrs := make(map[rune]int)
	total := 0
	count := 0
	for _, line := range a {
		if line == "" {
			total += calcUnique(ltrs, count)
			ltrs = make(map[rune]int)
			count = 0
		} else {
			for _, i := range line {
				ltrs[i] = ltrs[i] + 1
			}
			count++
		}
	}
	// Catch the last input
	total += calcUnique(ltrs, count)

	return strconv.Itoa(total)
}
func calcTotal(in map[rune]bool) int {
	total := 0
	for i := range in {
		if i >= 'a' && i <= 'z' {
			total++
		}
	}
	return total
}
func calcUnique(in map[rune]int, count int) int {
	total := 0
	for i, j := range in {
		if i >= 'a' && i <= 'z' && j == count {
			total++
		}
	}
	return total
}
