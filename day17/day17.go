package day17

import (
	"fmt"
)

type Range struct {
	min int
	max int
}

func (r *Range) updateMinMax(i int) {
	if i < r.min {
		r.min = i
	}
	if i > r.max {
		r.max = i
	}
}

type field struct {
	m    map[string]bool
	x    Range
	y    Range
	z    Range
	w    Range
	axes int
}

func (field *field) String() string {
	out := ""
	for l := field.w.min; l <= field.w.max; l++ {
		for k := field.z.min; k <= field.z.max; k++ {
			out += fmt.Sprintf("Z=%d-%d\n", k, l)
			for j := field.y.min; j <= field.y.max; j++ {
				for i := field.x.min; i <= field.x.max; i++ {
					if field.get(i, j, k, l) {
						out += "#"
					} else {
						out += "."
					}
				}
				out += fmt.Sprintf(" - %d\n", j)
			}
		}
	}
	return out
}
func (field *field) set(x int, y int, z int, w int) {
	loc := fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)
	field.m[loc] = true
	field.x.updateMinMax(x)
	field.y.updateMinMax(y)
	field.z.updateMinMax(z)
	if field.axes > 3 {
		field.w.updateMinMax(w)

	}
}
func (field *field) get(x int, y int, z int, w int) bool {
	loc := fmt.Sprintf("%d,%d,%d,%d", x, y, z, w)
	val, found := field.m[loc]
	return val && found
}
func (field *field) forEach(f func(*field, int, int, int, int)) {
	minX, maxX := field.x.min-1, field.x.max+1
	minY, maxY := field.y.min-1, field.y.max+1
	minZ, maxZ := field.z.min-1, field.z.max+1
	var minW, maxW int
	if field.axes > 3 {
		minW = field.w.min - 1
		maxW = field.w.max + 1
	} else {
		minW = 0
		maxW = 0
	}

	for i := minX; i <= maxX; i++ {
		for j := minY; j <= maxY; j++ {
			for k := minZ; k <= maxZ; k++ {
				for l := minW; l <= maxW; l++ {
					f(field, i, j, k, l)
				}
			}
		}
	}
}

func neighbors(f *field, x int, y int, z int, w int) int {
	total := 0
	count := 0
	var minW, maxW int
	if f.axes > 3 {
		minW = w - 1
		maxW = w + 1
	} else {
		minW = 0
		maxW = 0
	}
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				for l := minW; l <= maxW; l++ {
					if i == x && j == y && k == z && l == w {
						continue
					}
					if f.get(i, j, k, l) {
						total++
					}
					count++
				}
			}
		}
	}
	return total
}

// Day17a solves the puzzle and returns a string
func Day17a(input []string) string {
	cels := field{make(map[string]bool), Range{100, -100}, Range{100, -100}, Range{100, -100}, Range{100, -100}, 3}

	// Set initial state
	for j, row := range input {
		for i, cel := range row {
			if cel == '#' {
				cels.set(i, j, 0, 0)
			}
		}
	}

	fmt.Printf("\n-----\nInput:\n%s", cels.String())
	for count := 0; count < 6; count++ {
		newCels := field{make(map[string]bool), Range{100, -100}, Range{100, -100}, Range{100, -100}, Range{100, -100}, 3}
		f := func(f *field, i int, j int, k int, l int) {
			nearby := neighbors(f, i, j, k, l)
			found := f.get(i, j, k, l)
			if (found && nearby >= 2 && nearby <= 3) || (!found && nearby == 3) {
				newCels.set(i, j, k, l)
			}
		}
		cels.forEach(f)
		fmt.Printf("\nRound %d:\n%s\n", count+1, newCels.String())
		cels = newCels
	}

	return fmt.Sprintf("%d", len(cels.m))
}

// Day17b solves the puzzle and returns a string
func Day17b(input []string) string {
	cels := field{make(map[string]bool), Range{100, -100}, Range{100, -100}, Range{100, -100}, Range{100, -100}, 4}

	// Set initial state
	for j, row := range input {
		for i, cel := range row {
			if cel == '#' {
				cels.set(i, j, 0, 0)
			}
		}
	}

	fmt.Printf("\n-----\nInput:\n%s", cels.String())
	for count := 0; count < 6; count++ {
		newCels := field{make(map[string]bool), Range{100, -100}, Range{100, -100}, Range{100, -100}, Range{100, -100}, 4}
		f := func(f *field, i int, j int, k int, l int) {
			nearby := neighbors(f, i, j, k, l)
			found := f.get(i, j, k, l)
			if (found && nearby >= 2 && nearby <= 3) || (!found && nearby == 3) {
				newCels.set(i, j, k, l)
			}
		}
		cels.forEach(f)
		fmt.Printf("\nRound %d:\n%s\n", count+1, newCels.String())
		cels = newCels
	}

	return fmt.Sprintf("%d", len(cels.m))
}
