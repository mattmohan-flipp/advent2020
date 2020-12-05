package day4

import (
	"io/ioutil"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	file, _ := ioutil.ReadFile("testdata/day4a.txt")

	a := strings.Split(string(file), "\n")

	expected := "2"

	result := Day4a(a)

	if result != expected {
		t.Errorf("Day4a: Result %s doesn't match expected %s.", result, expected)
	}
}

func TestBValid(t *testing.T) {
	file, _ := ioutil.ReadFile("testdata/day4bv.txt")

	a := strings.Split(string(file), "\n")

	expected := "4"

	result := Day4b(a)

	if result != expected {
		t.Errorf("Day4b: Result %s doesn't match expected %s.", result, expected)
	}
}

func TestBInvalid(t *testing.T) {
	f := `eyr:1972 cid:100
	hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926
	
	iyr:2019
	hcl:#602927 eyr:1967 hgt:170cm
	ecl:grn pid:012533040 byr:1946
	
	hcl:dab227 iyr:2012
	ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277
	
	hgt:59cm ecl:zzz
	eyr:2038 hcl:74454a iyr:2023
	pid:3556412378 byr:2007`

	a := strings.Split(f, "\n")

	expected := "0"

	result := Day4b(a)

	if result != expected {
		t.Errorf("Day4b: Result %s doesn't match expected %s.", result, expected)
	}
}
