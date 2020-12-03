package day3

import "testing"

func TestA(t *testing.T) {
	a := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	expected := "7"

	result := Day3a(a)

	if result != expected {
		t.Errorf("Day3a: Result %s doesn't match expected %s.", result, expected)
	}
}

func TestB(t *testing.T) {
	a := []string{
		"..##.......",
		"#...#...#..",
		".#....#..#.",
		"..#.#...#.#",
		".#...##..#.",
		"..#.##.....",
		".#.#.#....#",
		".#........#",
		"#.##...#...",
		"#...##....#",
		".#..#...#.#",
	}

	expected := "336"

	result := Day3b(a)

	if result != expected {
		t.Errorf("Day3b: Result %s doesn't match expected %s.", result, expected)
	}
}
