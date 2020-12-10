package day9

import (
	"container/list"
	"strconv"
)

type OrderedSet struct {
	m      map[int]bool
	l      *list.List
	length int
}

func (o *OrderedSet) RemoveLast() {
	old := o.l.Remove(o.l.Front())
	delete(o.m, old.(int))
}

func (o *OrderedSet) Add(item int) {
	if o.l.Len() >= o.length {
		o.RemoveLast()
	}
	o.l.PushBack(item)
	o.m[item] = true
}

// Contains checks whether the item exists in the set
func (o *OrderedSet) Contains(item int) bool {
	_, found := o.m[item]
	return found
}

// Day9a solves the puzzle and returns a string
func Day9a(a []string) string {
	b := make([]int, len(a))
	for i, j := range a {
		b[i], _ = strconv.Atoi(j)
	}
	return strconv.Itoa(SolveA(b, 25))
}

// SolveA comment
func SolveA(inputs []int, window int) int {
	set := OrderedSet{make(map[int]bool), list.New(), window}
	for i, input := range inputs {
		if i > window {
			paired := false
			for k := range set.m {
				if set.Contains(input - k) {
					paired = true
					break
				}
			}
			if !paired {
				return input
			}
		}
		set.Add(input)
	}
	return -1
}

// Day9b solves the puzzle and returns a string
func Day9b(a []string) string {
	b := make([]int, len(a))
	for i, j := range a {
		b[i], _ = strconv.Atoi(j)
	}
	return SolveB(b, 25)
}

// SolveA comment
func SolveB(inputs []int, window int) string {
	target := SolveA(inputs, window)
	for i, curStart := range inputs {
		total := curStart
		min := curStart
		max := curStart
		for _, val := range inputs[i+1:] {
			total += val
			if val > max {
				max = val
			}
			if val < min {
				min = val
			}
			if total == target {
				return strconv.Itoa(min + max)
			}
			if total > target {
				break
			}
		}
	}
	return "Unknown"
}
