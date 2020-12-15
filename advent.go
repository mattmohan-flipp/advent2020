package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"mattmohan.com/advent2020/day1"
	"mattmohan.com/advent2020/day10"
	"mattmohan.com/advent2020/day11"
	"mattmohan.com/advent2020/day12"
	"mattmohan.com/advent2020/day13"
	"mattmohan.com/advent2020/day14"
	"mattmohan.com/advent2020/day15"
	"mattmohan.com/advent2020/day2"
	"mattmohan.com/advent2020/day3"
	"mattmohan.com/advent2020/day4"
	"mattmohan.com/advent2020/day5"
	"mattmohan.com/advent2020/day6"
	"mattmohan.com/advent2020/day7"
	"mattmohan.com/advent2020/day8"
	"mattmohan.com/advent2020/day9"
)

// DailyTask defines each puzzle task
type DailyTask func([]string) string

func main() {
	day := flag.String("day", "3a", "Which day+variant to run (e.g. 1a)")
	file := flag.String("file", "-", "The input file (- for stdin)")
	flag.Parse()

	taskname := "day" + *day
	task := GetFunc(taskname)

	inputFile := GetFile(*file)
	splitInput := ReadFileToStringArray(inputFile)

	result := task(splitInput)

	println("Got: ", result)
}

// ReadFileToStringArray takes a file handle and returns a slice containing each line as a string
func ReadFileToStringArray(inputFile *os.File) []string {
	inputBytes, err := ioutil.ReadAll(inputFile)
	if err != nil {
		println("Error reading data: ", err.Error())
		os.Exit(1)
	}
	input := string(inputBytes)
	return strings.Split(input, "\n")
}

// GetFunc takes a task name and returns the task function
func GetFunc(taskname string) DailyTask {
	functions := map[string]DailyTask{
		"day1a":  day1.Day1a,
		"day1b":  day1.Day1b,
		"day2a":  day2.Day2a,
		"day2b":  day2.Day2b,
		"day3a":  day3.Day3a,
		"day3b":  day3.Day3b,
		"day4a":  day4.Day4a,
		"day4b":  day4.Day4b,
		"day5a":  day5.Day5a,
		"day5b":  day5.Day5b,
		"day6a":  day6.Day6a,
		"day6b":  day6.Day6b,
		"day7a":  day7.Day7a,
		"day7b":  day7.Day7b,
		"day8a":  day8.Day8a,
		"day8b":  day8.Day8b,
		"day9a":  day9.Day9a,
		"day9b":  day9.Day9b,
		"day10a": day10.Day10a,
		"day10b": day10.Day10b,
		"day11a": day11.Day11a,
		"day11b": day11.Day11b,
		"day12a": day12.Day12a,
		"day12b": day12.Day12b,
		"day13a": day13.Day13a,
		"day13b": day13.Day13b,
		"day14a": day14.Day14a,
		"day14b": day14.Day14b,
		"day15a": day15.Day15a,
		"day15b": day15.Day15b,
	}

	task, found := functions[taskname]
	if !found {
		println("Invalid day/variant provided", taskname)
		os.Exit(1)
	}

	return task
}

// GetFile takes the file param and returns an open file handle
func GetFile(file string) *os.File {
	if file != "-" {
		openFile, err := os.Open(file)
		if err != nil {
			println("Error opening file: ", err.Error())
			os.Exit(1)
		}
		return openFile
	}
	return os.Stdin
}
