package day4

import "testing"

func TestA(t *testing.T) {
	t.Skip("Not implemented")

	a := []string{
		"",
		"",
		"",
	}

	expected := ""

	result := Day4a(a)

	if result != expected {
		t.Errorf("Day4a: Result %s doesn't match expected %s.", result, expected)
	}
}

func TestB(t *testing.T) {
	t.Skip("Not implemented")

	a := []string{"",
		"",
		""}

	expected := ""

	result := Day4b(a)

	if result != expected {
		t.Errorf("Day4b: Result %s doesn't match expected %s.", result, expected)
	}
}
