package day11

import (
	"strconv"

	"mattmohan.com/advent2020/intutils"
)

const (
	floor = iota
	empty
	full
)

// Grid is the current state of the grid
type Grid [][]int

func (grid Grid) countOccupied() int {
	occ := 0
	for _, i := range grid {
		for _, j := range i {
			if j == full {
				occ++
			}
		}
	}
	return occ
}

func convertInputToGrid(a []string) Grid {
	b := make(Grid, len(a))
	for i := range a {
		if a[i] == "" {
			continue
		}
		b[i] = make([]int, len(a[i]))
		for j, curr := range a[i] {
			if curr == '#' {
				b[i][j] = full
			} else if curr == 'L' {
				b[i][j] = empty
			} else {
				b[i][j] = floor
			}
		}
	}
	return b
}

// Ray holds the slope of a direction
type Ray struct {
	x int
	y int
}

func (ray Ray) calcNextRayPosition(row int, col int) (int, int) {
	return row + ray.y, col + ray.x
}
func checkInbounds(i int, j int, inp [][]int) bool {
	maxI := len(inp)
	if i < 0 || i >= maxI {
		return false
	}
	if j < 0 || j >= len(inp[i]) {
		return false
	}
	return true
}

func calcFar(inp Grid, row int, col int) int {
	nearby := 0
	rays := []Ray{
		{0, -1},
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
	}

	for _, ray := range rays {
		for i, j := ray.calcNextRayPosition(row, col); checkInbounds(i, j, inp); i, j = ray.calcNextRayPosition(i, j) {

			if inp[i][j] == full {
				nearby++
				break
			} else if inp[i][j] == empty {
				break
			}
		}
	}

	return nearby
}

func calcNearby(inp Grid, row int, col int) int {
	nearby := 0
	startRow := intutils.Max(row-1, 0)
	endRow := intutils.Min(row+1, len(inp)-1)

	for i := startRow; i <= endRow; i++ {
		startCol := intutils.Max(col-1, 0)
		endCol := intutils.Min(col+1, len(inp[i])-1)

		for j := startCol; j <= endCol; j++ {
			if i == row && j == col {
				continue
			}
			if inp[i][j] == full {
				nearby++
			}
		}

	}

	return nearby
}

func calcNext(inp Grid, threshold int, far bool) (Grid, int) {
	changes := 0
	out := make(Grid, len(inp))
	for i := range inp {
		out[i] = make([]int, len(inp[i]))
		for j := range inp[i] {
			var neighbors int
			if far {
				neighbors = calcFar(inp, i, j)
			} else {
				neighbors = calcNearby(inp, i, j)
			}
			if inp[i][j] == full && neighbors >= threshold {
				out[i][j] = empty
				changes++
			} else if inp[i][j] == empty && neighbors == 0 {
				out[i][j] = full
				changes++
			} else {
				out[i][j] = inp[i][j]
			}
		}
	}
	return out, changes
}

// Day11a solves the puzzle and returns a string
func Day11a(a []string) string {
	grid := convertInputToGrid(a)
	changes := 0
	for grid, changes = calcNext(grid, 4, false); changes != 0; grid, changes = calcNext(grid, 4, false) {
	}

	occupied := grid.countOccupied()

	return strconv.Itoa(occupied)
}

// Day11b solves the puzzle and returns a string
func Day11b(a []string) string {
	grid := convertInputToGrid(a)
	changes := 0

	count := 0

	for grid, changes = calcNext(grid, 5, true); changes != 0; grid, changes = calcNext(grid, 5, true) {
		count++
	}

	occupied := grid.countOccupied()

	return strconv.Itoa(occupied)
}
