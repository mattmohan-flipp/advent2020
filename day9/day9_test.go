package day9

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := 127
	intInput := make([]int, len(input))
	for i, j := range input {
		conv, _ := strconv.Atoi(j)
		intInput[i] = conv
	}

	results := SolveA(intInput, 5)
	if results != expected {
		t.Errorf("Day9a: Result %d doesn't match expected %d.", results, expected)
	}
}

func TestB(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "62"
	intInput := make([]int, len(input))
	for i, j := range input {
		conv, _ := strconv.Atoi(j)
		intInput[i] = conv
	}
	results := SolveB(intInput, 5)
	if results != expected {
		t.Errorf("Day9b: Result %s doesn't match expected %s.", results, expected)
	}
}
