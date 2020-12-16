package day16

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type class struct {
	name      string
	lowStart  int
	lowEnd    int
	highStart int
	highEnd   int
}

var classRegex = regexp.MustCompile(`([^:]+): ([\d]+)-([\d]+) or ([\d]+)-([\d]+)`)

func getClass(input string) class {
	parts := classRegex.FindStringSubmatch(input)
	name := parts[1]
	lowStart, _ := strconv.Atoi(parts[2])
	lowEnd, _ := strconv.Atoi(parts[3])
	highStart, _ := strconv.Atoi(parts[4])
	highEnd, _ := strconv.Atoi(parts[5])
	cls := class{name, lowStart, lowEnd, highStart, highEnd}
	return cls
}
func getTicket(input string, valid []bool) ([]int, bool) {
	cols := strings.Split(input, ",")
	var colNums []int
	ok := true
	for _, i := range cols {
		num, _ := strconv.Atoi(i)
		if !valid[num] {
			ok = false
		}
		colNums = append(colNums, num)
	}
	return colNums, ok
}

func inClass(nums []int, class class) bool {
	ok := true
	for _, j := range nums {
		if !(j >= class.highStart && j <= class.highEnd) && !(j >= class.lowStart && j <= class.lowEnd) {
			ok = false
			break
		}
	}
	return ok
}

// Day16a solves the puzzle and returns a string
func Day16a(input []string) string {
	state := 0
	valid := make([]bool, 1000)
	invalid := 0

	for i := range valid {
		valid[i] = false
	}
	for _, i := range input {
		if i == "" {
			state++
			continue
		}
		if i == "nearby tickets:" || i == "your ticket:" {
			continue
		}
		if state == 0 {
			class := getClass(i)
			for i := class.lowStart; i <= class.lowEnd; i++ {
				valid[i] = true
			}
			for i := class.highStart; i <= class.highEnd; i++ {
				valid[i] = true
			}
		} else if state == 2 {
			values, _ := getTicket(i, valid)
			fmt.Printf("%s -> %v\n", i, values)
			for _, j := range values {
				if !valid[j] {
					invalid += j
				}
			}
		}
	}

	return strconv.Itoa(invalid)
}

// Day16b solves the puzzle and returns a string
func Day16b(input []string) string {
	state := 0
	valid := make([]bool, 1000)
	var validTickets [][]int
	var myTicket []int
	var classes []class
	// Map class names back to columns
	classMatches := make(map[string][]int)
	classMap := make(map[string]int)

	for i := range valid {
		valid[i] = false
	}
	for _, i := range input {
		if i == "" {
			state++
			continue
		}
		if state == 0 {
			class := getClass(i)
			for i := class.lowStart; i <= class.lowEnd; i++ {
				valid[i] = true
			}
			for i := class.highStart; i <= class.highEnd; i++ {
				valid[i] = true
			}
			classes = append(classes, class)
		} else if state == 1 {
			myTicket, _ = getTicket(i, valid)
		} else if state == 2 {
			values, ok := getTicket(i, valid)
			if ok {
				validTickets = append(validTickets, values)
			}
		}
	}

	width := len(validTickets[0])
	col := make([][]int, width)
	for i := 0; i < width; i++ {
		for _, j := range validTickets {
			col[i] = append(col[i], j[i])
		}

		for _, class := range classes {
			ok := inClass(col[i], class)
			if ok {
				classMatches[class.name] = append(classMatches[class.name], i)
			}
		}
	}
	for len(classMatches) > 0 {
		for i, j := range classMatches {
			if len(j) == 1 {
				classMap[i] = j[0]
				delete(classMatches, i)
				for k := range classMatches {
					for l := range classMatches[k] {
						if classMatches[k][l] == j[0] {
							classMatches[k] = append(classMatches[k][0:l], classMatches[k][l+1:]...)
							break
						}
					}
				}
			}
		}
	}
	total := 1
	for name := range classMap {
		if strings.HasPrefix(name, "departure") {
			total *= myTicket[classMap[name]]
		}
	}
	return fmt.Sprintf("%d", total)
}
