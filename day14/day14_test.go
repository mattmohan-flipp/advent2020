package day14

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "165"

	results := Day14a(input)
	if results != expected {
		t.Errorf("Day14a: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_b.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "208"

	results := Day14b(input)
	if results != expected {
		t.Errorf("Day14b: Result %s doesn't match expected %s.", results, expected)
	}
}
