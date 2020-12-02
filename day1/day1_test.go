package day1

import "testing"

func TestA(t *testing.T) {
	a := []string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	}
	expected := "514579"

	result := Day1a(a)

	if result != expected {
		t.Errorf("Day1a: Result %s doesn't match expected %s.", result, expected)
	}
}
func TestB(t *testing.T) {
	a := []string{
		"1721",
		"979",
		"366",
		"299",
		"675",
		"1456",
	}
	expected := "241861950"

	result := Day1b(a)

	if result != expected {
		t.Errorf("Day1b: Result %s doesn't match expected %s.", result, expected)
	}
}
