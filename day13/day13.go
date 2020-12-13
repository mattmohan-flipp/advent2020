package day13

import (
	"container/list"
	"math"
	"strconv"
	"strings"
)

// Bus is a bus
type Bus struct {
	position int
	step     int
}

// Day13a solves the puzzle and returns a string
func Day13a(input []string) string {
	timestamp, _ := strconv.Atoi(input[0])
	buses := make([]int, 0)
	for _, bus := range strings.Split(input[1], ",") {
		busNum, err := strconv.Atoi(bus)
		if err == nil {
			buses = append(buses, busNum)
		}
	}
	bestBus := 0
	wait := math.MaxInt64
	for _, bus := range buses {
		curWait := bus - timestamp%bus
		if curWait < wait {
			wait = curWait
			bestBus = bus
		}
	}

	return strconv.Itoa(bestBus * wait)
}

// Day13b solves the puzzle and returns a string
func Day13b(input []string) string {
	buses := list.New()
	for i, bus := range strings.Split(input[1], ",") {
		busNum, err := strconv.Atoi(bus)
		if err == nil {
			buses.PushBack(Bus{i, busNum})
		}
	}

	step := buses.Front().Value.(Bus).step
	t := step
	for e := buses.Front().Next(); e != nil; e = e.Next() {
		for bus := e.Value.(Bus); (t+bus.position)%bus.step != 0; t += step {
		}
		step *= e.Value.(Bus).step
	}
	return strconv.Itoa(t)
}
