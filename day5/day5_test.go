package day5

import (
	"strconv"
	"testing"
)

func TestDecode(t *testing.T) {
	inputs := []string{
		"BFFFBBFRRR",
		"FFFBBBFRRR",
		"BBFFBBFRLL",
	}
	expected := []string{
		"567",
		"119",
		"820",
	}

	results := []int{
		DecodeBoardingPass(inputs[0]),
		DecodeBoardingPass(inputs[1]),
		DecodeBoardingPass(inputs[2]),
	}
	for i := range inputs {
		if strconv.Itoa(results[i]) != expected[i] {
			t.Errorf("Day4a: Result%d %d doesn't match expected %s.", i, results[i], expected[i])
		}
	}
}
