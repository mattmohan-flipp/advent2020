package day17

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "112"

	results := Day17a(input)
	if results != expected {
		t.Errorf("Day17a: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "1"

	results := Day17b(input)
	if results != expected {
		t.Errorf("Day17b: Result %s doesn't match expected %s.", results, expected)
	}
}
