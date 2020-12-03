package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"

	"mattmohan.com/advent2020/day1"
	"mattmohan.com/advent2020/day2"
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
		"day1a": day1.Day1a,
		"day1b": day1.Day1b,
		"day2a": day2.Day2a,
		"day2b": day2.Day2b,
	}

	task, found := functions[taskname]
	if !found {
		println("Invalid day/variant provided")
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
	} else {
		return os.Stdin
	}
}
