package day5

import (
	"strconv"
)

// Day5a solves the puzzle and returns a string
func Day5a(a []string) string {
	highest := 0
	for _, i := range a {
		cur := DecodeBoardingPass(i)
		if cur > highest {
			highest = cur
		}
	}
	return strconv.Itoa(highest)
}

// MaxUint is the max unsigned int
const MaxUint = ^uint(0)

// MaxInt is the max signed int
const MaxInt = int(MaxUint >> 1)

// Day5b solves the puzzle and returns a string
func Day5b(a []string) string {
	seats := make(map[int]bool)
	min, max := MaxInt, 0
	for _, i := range a {
		seat := DecodeBoardingPass(i)
		seats[seat] = true
		if seat > max {
			max = seat
		}
		if seat < min {
			min = seat
		}
	}

	for i := min; i < max; i++ {
		_, found := seats[i]
		if !found && seats[i-1] && seats[i+1] {
			return strconv.Itoa(i)
		}
	}
	return ""
}

// DecodeBoardingPass turns the string ito the ID
func DecodeBoardingPass(pass string) int {
	rowS := pass[0:7]
	colS := pass[7:10]

	row := 0
	col := 0
	for _, i := range rowS {
		row = row << 1
		if i == 'B' {
			row = row | 1
		}
	}
	for _, i := range colS {
		col = col << 1
		if i == 'R' {
			col = col | 1
		}
	}

	return 8*row + col
}
