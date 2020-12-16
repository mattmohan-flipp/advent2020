package day16

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "71"

	results := Day16a(input)
	if results != expected {
		t.Errorf("Day16a: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "1"

	results := Day16b(input)
	if results != expected {
		t.Errorf("Day16b: Result %s doesn't match expected %s.", results, expected)
	}
}
