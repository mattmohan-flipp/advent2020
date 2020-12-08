package day8

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "5"

	results := Day8a(input)
	if results != expected {
		t.Errorf("Day7a: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "8"

	results := Day8b(input)
	if results != expected {
		t.Errorf("Day8b: Result %s doesn't match expected %s.", results, expected)
	}
}
