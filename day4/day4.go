package day4

import (
	"regexp"
	"strconv"
	"strings"
)

// Validations for part B; also, the keys are the set of required fields for part A
var regexs = map[string]*regexp.Regexp{
	"byr": regexp.MustCompile(`^(19[2-9][0-9])|(200[0-2])$`),
	"iyr": regexp.MustCompile(`^201[0-9]|2020$`),
	"eyr": regexp.MustCompile(`^202[0-9]|2030$`),
	"hgt": regexp.MustCompile(`^1[5-8][0-9]cm|19[0-3]cm|59in|6[0-9]in|7[0-6]in$`),
	"hcl": regexp.MustCompile(`^#[0-9a-f]{6}$`),
	"ecl": regexp.MustCompile(`^amb|blu|brn|gry|grn|hzl|oth$`),
	"pid": regexp.MustCompile(`^[0-9]{9}$`),
}

// Day4a solves the puzzle and returns a string
func Day4a(a []string) string {
	return genericMatch(a, func(tkn string, val string) bool {
		return val != ""
	})
}

// Day4b solves the puzzle and returns a string
func Day4b(a []string) string {
	return genericMatch(a, func(tkn string, val string) bool {
		return regexs[tkn].MatchString(val)
	})
}

func genericMatch(a []string, comparator func(tkn string, val string) bool) string {
	total := 0
	token := make(map[string]string)

	// Ugly hack to ensure the last passport gets scanned
	withNewline := append(a, "")

	for _, i := range withNewline {
		if i == "" {
			match := true
			for j := range regexs {
				if !comparator(j, token[j]) {
					match = false
					break
				}
			}
			if match {
				total++
			}
			token = make(map[string]string)
		} else {
			line := strings.Split(i, " ")
			for _, j := range line {
				split := strings.Split(j, ":")
				if len(split) == 2 {
					token[split[0]] = split[1]
				} else {
					println("Couldn't process: ", j)
				}
			}
		}
	}
	return strconv.Itoa(total)
}
