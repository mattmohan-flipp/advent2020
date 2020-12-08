package day8

import (
	"strconv"
	"strings"
)

// Day8a solves the puzzle and returns a string
func Day8a(a []string) string {
	acc, _ := evalProg(a)
	return strconv.Itoa(acc)
}

// Day8b solves the puzzle and returns a string
func Day8b(a []string) string {
	for i, orig := range a {
		if a[i][0:3] == "acc" {
			continue
		}

		if orig[0:3] == "nop" {
			a[i] = strings.Replace(orig, "nop", "jmp", 1)
		} else {
			a[i] = strings.Replace(orig, "jmp", "nop", 1)
		}

		acc, found := evalProg(a)
		if found {
			return strconv.Itoa(acc)
		}
		a[i] = orig
	}

	return "Not Found"
}

func evalProg(a []string) (int, bool) {
	acc := 0
	steps := make(map[int]bool)
	for cur := 0; cur < len(a); {
		if steps[cur] {
			return acc, false
		}
		steps[cur] = true
		inst := a[cur][0:3]
		value, _ := strconv.Atoi(a[cur][4:])

		if inst == "jmp" {
			cur += value
		} else if inst == "acc" {
			acc += value
			cur++
		} else {
			cur++
		}
	}
	return acc, true
}
