package day1

import (
	"fmt"
	"strconv"
)

type point struct {
	x int
	y int
}

func strings2map(s []string) map[int]bool {
	out := make(map[int]bool)
	for _, i := range s {
		in, _ := strconv.Atoi(i)
		out[in] = true
	}
	return out
}

// Day1a turns input strings into an output string
func Day1a(s []string) string {
	a := strings2map(s)

	for i := range a {
		_, found := a[2020-i]
		if found {
			j := 2020 - i
			return fmt.Sprint(i * j)
		}
	}
	return "N/A"
}

// Day1b turns input strings into an output string
func Day1b(s []string) string {
	a := strings2map(s)

	pairs := make(map[int]point)
	for i := range a {
		// i+j+k=2020 => 2020-i = j+k
		// So lookup 2020-a in a set containing all previous j+k pairings
		val, found := pairs[(2020 - i)]
		if found {
			return fmt.Sprint(i * val.x * val.y)
		}

		// Precompute potential pairs
		for j := range a {
			pairs[i+j] = point{i, j}
		}
		a[i] = true
	}
	return "N/A"
}
