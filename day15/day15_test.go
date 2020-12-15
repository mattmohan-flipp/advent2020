package day15

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "436"

	results := Day15a(input)
	if results != expected {
		t.Errorf("Day15a: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestA2(t *testing.T) {
	expected := "1"
	results := Day15a([]string{"1,3,2"})
	if results != expected {
		t.Errorf("Day15a: Result %s doesn't match expected %s.", results, expected)
	}
	expected = "10"
	results = Day15a([]string{"2,1,3"})
	if results != expected {
		t.Errorf("Day15a: Result %s doesn't match expected %s.", results, expected)
	}

	expected = "27"
	results = Day15a([]string{"1,2,3"})
	if results != expected {
		t.Errorf("Day15a: Result %s doesn't match expected %s.", results, expected)
	}

	expected = "78"
	results = Day15a([]string{"2,3,1"})
	if results != expected {
		t.Errorf("Day15a: Result %s doesn't match expected %s.", results, expected)
	}

	expected = "438"
	results = Day15a([]string{"3,2,1"})
	if results != expected {
		t.Errorf("Day15a: Result %s doesn't match expected %s.", results, expected)
	}

	expected = "1836"
	results = Day15a([]string{"3,1,2"})
	if results != expected {
		t.Errorf("Day15a: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "175594"

	results := Day15b(input)
	if results != expected {
		t.Errorf("Day15b: Result %s doesn't match expected %s.", results, expected)
	}
}
