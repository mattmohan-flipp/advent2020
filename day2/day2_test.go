package day2

import "testing"

func TestA(t *testing.T) {
	a := []string{"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc"}

	expected := "2"

	result := Day2a(a)

	if result != expected {
		t.Errorf("Day2a: Result %s doesn't match expected %s.", result, expected)
	}
}

func TestB(t *testing.T) {
	a := []string{"1-3 a: abcde",
		"1-3 b: cdefg",
		"2-9 c: ccccccccc"}

	expected := "1"

	result := Day2b(a)

	if result != expected {
		t.Errorf("Day2a: Result %s doesn't match expected %s.", result, expected)
	}
}
