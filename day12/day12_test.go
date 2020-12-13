package day12

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "25"

	results := Day12a(input)
	if results != expected {
		t.Errorf("Day12a: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "286"

	results := Day12b(input)
	if results != expected {
		t.Errorf("Day12b: Result %s doesn't match expected %s.", results, expected)
	}
}
