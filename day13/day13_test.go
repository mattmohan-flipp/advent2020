package day13

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "295"

	results := Day13a(input)
	if results != expected {
		t.Errorf("Day13a: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "1068781"

	results := Day13b(input)
	if results != expected {
		t.Errorf("Day13b: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB2(t *testing.T) {
	input := []string{"", "17,x,13,19"}
	expected := "3417"

	results := Day13b(input)
	if results != expected {
		t.Errorf("Day13b: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB3(t *testing.T) {
	input := []string{"", "67,7,59,61"}
	expected := "754018"

	results := Day13b(input)
	if results != expected {
		t.Errorf("Day13b: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB4(t *testing.T) {
	input := []string{"", "67,x,7,59,61"}
	expected := "779210"

	results := Day13b(input)
	if results != expected {
		t.Errorf("Day13b: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB5(t *testing.T) {
	input := []string{"", "67,7,x,59,61"}
	expected := "1261476"

	results := Day13b(input)
	if results != expected {
		t.Errorf("Day13b: Result %s doesn't match expected %s.", results, expected)
	}
}

func TestB6(t *testing.T) {
	input := []string{"", "1789,37,47,1889"}
	expected := "1202161486"

	results := Day13b(input)
	if results != expected {
		t.Errorf("Day13b: Result %s doesn't match expected %s.", results, expected)
	}
}
