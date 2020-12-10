package day10

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := 220
	intInput := make([]int, len(input))
	for i, j := range input {
		conv, _ := strconv.Atoi(j)
		intInput[i] = conv
	}

	results := SolveA(intInput)
	if results != expected {
		t.Errorf("Day10a: Result %d doesn't match expected %d.", results, expected)
	}
}

func TestB(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := 19208
	intInput := make([]int, len(input))
	for i, j := range input {
		conv, _ := strconv.Atoi(j)
		intInput[i] = conv
	}
	results := SolveB(intInput)
	if results != expected {
		t.Errorf("Day10b: Result %d doesn't match expected %d.", results, expected)
	}
}
