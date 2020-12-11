package day11

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "37"

	results := Day11a(input)
	if results != expected {
		t.Errorf("Day10a: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "26"

	results := Day11b(input)
	if results != expected {
		t.Errorf("Day10b: Result %s doesn't match expected %s.", results, expected)
	}
}
