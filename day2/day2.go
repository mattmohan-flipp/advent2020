package day2

import (
	"strconv"
	"strings"
)

// Day2a solves the puzzle and returns a string
func Day2a(a []string) string {
	total := 0
	for _, i := range a {
		split := strings.Split(i, " ")
		minMax := strings.Split(split[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		match := strings.Split(split[1], ":")[0]
		comp := split[2]
		count := 0
		for _, j := range comp {
			if strings.ContainsRune(match, j) {
				count++
			}
		}
		if count >= min && count <= max {
			total++
		}
	}
	return strconv.Itoa(total)
}

func Day2b(a []string) string {
	total := 0
	for _, i := range a {
		split := strings.Split(i, " ")
		minMax := strings.Split(split[0], "-")
		min, _ := strconv.Atoi(minMax[0])
		max, _ := strconv.Atoi(minMax[1])
		match := strings.Split(split[1], ":")[0]
		comp := split[2]
		lowMatch := match == string(comp[min-1])
		highMatch := match == string(comp[max-1])
		if (lowMatch || highMatch) && !(lowMatch && highMatch) {
			total++
		}
	}
	return strconv.Itoa(total)
}
