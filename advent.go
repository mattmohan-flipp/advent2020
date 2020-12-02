package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"mattmohan.com/advent2020/day1"
	"mattmohan.com/advent2020/day2"
)

// DailyTask defines each puzzle task
type DailyTask func([]string) string

func main() {
	functions := map[string]DailyTask{
		"day1a": day1.Day1a,
		"day1b": day1.Day1b,
		"day2a": day2.Day2a,
		"day2b": day2.Day2b,
	}

	day := flag.Int("day", 1, "Which day to run")
	variant := flag.String("variant", "a", "Which variant to run (a/b)")
	file := flag.String("file", "-", "The input file")

	flag.Parse()

	taskname := "day" + strconv.Itoa(*day) + *variant
	task, found := functions[taskname]
	if !found {
		println("Invalid day/variant provided")
		os.Exit(1)
	}

	inputFile := os.Stdin
	if *file != "-" {
		openFile, err := os.Open(*file)
		if err != nil {
			println("Error opening file: ", err.Error())
			os.Exit(1)
		}
		inputFile = openFile
	}

	inputBytes, err := ioutil.ReadAll(inputFile)
	if err != nil {
		println("Error reading data: ", err.Error())
		os.Exit(1)
	}
	input := string(inputBytes)
	splitInput := strings.Split(input, "\n")
	result := task(splitInput)
	println("Got: ", result)
}
