package day10

import (
	"sort"
	"strconv"
)

// Day10a solves the puzzle and returns a string
func Day10a(a []string) string {
	b := make([]int, len(a))
	for i, j := range a {
		tmp, _ := strconv.Atoi(j)
		b[i] = tmp
	}
	return strconv.Itoa(SolveA(b))
}

// SolveA comment
func SolveA(inputs []int) int {
	count1, count3 := 0, 1
	last := 0
	sort.Ints(inputs)

	for _, j := range inputs {
		tmp := j - last
		if tmp == 1 {
			count1++
		} else if tmp == 3 {
			count3++
		}
		last = j
	}

	return count1 * count3
}

// Day10b solves the puzzle and returns a string
func Day10b(a []string) string {
	b := make([]int, len(a))
	for i, j := range a {
		tmp, _ := strconv.Atoi(j)
		b[i] = tmp
	}
	return strconv.Itoa(SolveB(b))
}

// SolveB comment
func SolveB(inputs []int) int {
	count := len(inputs)
	costs := make(map[int]int, count+2)
	sort.Ints(inputs)

	// Insert the static endpoint
	costs[inputs[count-1]+3] = 1

	for i := (count - 1); i >= 0; i-- {
		cur := inputs[i]
		costs[cur] = costs[cur+1] + costs[cur+2] + costs[cur+3]
	}

	return costs[1] + costs[2] + costs[3]
}
