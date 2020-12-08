package day7

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	inputs, _ := ioutil.ReadFile("testdata/input_a.txt")
	input := strings.Split(string(inputs), "\n")
	expected := "4"

	results := Day7a(input)
	if results != expected {
		t.Errorf("Day7a: Result %s doesn't match expected %s.", results, expected)
	}
}
