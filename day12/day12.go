package day12

import (
	"fmt"
	"math"
	"strconv"

	"mattmohan.com/advent2020/intutils"
)

// Ship represents the position and direction of a ship or waypoint at sea
type Ship struct {
	x   int
	y   int
	dir int
}

func (ship Ship) rotateWaypoint(waypoint Ship, angle int) Ship {
	sin, cos := math.Sincos(float64(angle) * math.Pi / 180.0)
	x := float64(waypoint.x)
	y := float64(waypoint.y)
	newX := int(math.Round(x*cos - y*sin))
	newY := int(math.Round(x*sin + y*cos))
	return Ship{newX, newY, 0}
}

func (ship Ship) String() string {
	return fmt.Sprintf("%d, %d, %d", ship.x, ship.y, ship.dir)
}

// Day12a solves the puzzle and returns a string
func Day12a(instructions []string) string {
	ship := Ship{0, 0, 0}
	for _, instruction := range instructions {
		num, _ := strconv.Atoi(instruction[1:])
		switch instruction[0] {
		case 'R':
			ship.dir -= num
			break
		case 'L':
			ship.dir += num
			break
		case 'N':
			ship.y += num
			break
		case 'S':
			ship.y -= num
			break
		case 'E':
			ship.x += num
			break
		case 'W':
			ship.x -= num
			break
		case 'F':
			sin, cos := math.Sincos(float64(ship.dir) / 180.0 * math.Pi)
			ship.x += int(cos * float64(num))
			ship.y += int(sin * float64(num))
			break
		}

	}

	x, y := intutils.Abs(ship.x), intutils.Abs(ship.y)

	return strconv.Itoa(x + y)
}

// Day12b solves the puzzle and returns a string
func Day12b(instructions []string) string {
	waypoint := Ship{10, 1, 0}
	ship := Ship{0, 0, 0}

	for _, instruction := range instructions {
		num, _ := strconv.Atoi(instruction[1:])

		switch instruction[0] {
		case 'R':
			waypoint = ship.rotateWaypoint(waypoint, -num)
			break
		case 'L':
			waypoint = ship.rotateWaypoint(waypoint, num)
			break
		case 'N':
			waypoint.y += num
			break
		case 'S':
			waypoint.y -= num
			break
		case 'E':
			waypoint.x += num
			break
		case 'W':
			waypoint.x -= num
			break
		case 'F':
			ship.x += waypoint.x * num
			ship.y += waypoint.y * num
			break
		}
	}

	x, y := intutils.Abs(ship.x), intutils.Abs(ship.y)

	return strconv.Itoa(x + y)
}
