package day3

import (
	"strconv"
)

type field struct {
	width  int
	height int
	a      []string
}

type run struct {
	field field
	xStep int
	yStep int
}

// Day3a solves the puzzle and returns a string
func Day3a(a []string) string {
	x := 0
	width := len(a[0])
	total := 0
	for _, j := range a {
		if j[x] == '#' {
			total++
		}

		x = (x + 3) % width
	}
	return strconv.Itoa(total)
}

// Day3b solves the puzzle and returns a string
func Day3b(a []string) string {
	xSteps := []int{1, 3, 5, 7, 1}
	ySteps := []int{1, 1, 1, 1, 2}

	width := len(a[0])
	height := len(a)
	c := make(chan int, len(xSteps))

	for i := range xSteps {
		// Making this parallel is actually slower (setup time outweighs savings) but...
		go func(xStep int, yStep int) {
			curTotal := 0
			for x, y := 0, 0; y < height; x, y = ((x + xStep) % width), (y + yStep) {
				if a[y][x] == '#' {
					curTotal++
				}
			}
			c <- curTotal
		}(xSteps[i], ySteps[i])
	}

	total := 1
	for range xSteps {
		total *= <-c
	}

	return strconv.Itoa(total)
}
