package day6

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "11"

	results := Day6a(input)
	if results != expected {
		t.Errorf("Day4a: Result %s doesn't match expected %s.", results, expected)
	}
}
