package day15

import (
	"fmt"
	"strconv"
	"strings"
)

func solve(input []string, count int) string {
	previous := make(map[int]int)
	curTurn := 0
	nums := strings.Split(input[0], ",")
	var lastSaid int
	for i, numString := range nums {
		num, _ := strconv.Atoi(numString)
		lastSaid = num
		if i < len(nums)-1 {
			previous[num] = curTurn
		}
		curTurn++
	}
	for ; curTurn < count; curTurn++ {
		lastSeen, found := previous[lastSaid]
		if found {
			previous[lastSaid] = curTurn - 1
			lastSaid = curTurn - lastSeen - 1
		} else {
			previous[lastSaid] = curTurn - 1
			lastSaid = 0
		}
	}
	return fmt.Sprintf("%d", lastSaid)
}

// Day15a solves the puzzle and returns a string
func Day15a(input []string) string {
	return solve(input, 2020)
}

// Day15b solves the puzzle and returns a string
func Day15b(input []string) string {
	return solve(input, 30000000)
}
