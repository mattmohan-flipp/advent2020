package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var memRegex = regexp.MustCompile(`(\d+)`)

func binStringFromDec(in string) string {
	val, _ := strconv.Atoi(in)
	return fmt.Sprintf("%036b", val)
}

func sumMem(mem map[int]string) int64 {
	var total int64
	total = 0
	for _, i := range mem {
		val, _ := strconv.ParseInt(i, 2, 64)
		println(val)
		total += val
	}
	return total
}

// Day14a solves the puzzle and returns a string
func Day14a(input []string) string {
	mem := make(map[int]string)
	mask := strings.Repeat("X", 36)
	for _, instruction := range input {
		parts := strings.Split(instruction, " = ")
		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			addr, _ := strconv.Atoi(memRegex.FindString(parts[0]))
			bits := []byte(binStringFromDec(parts[1]))

			for i, bit := range mask {
				if bit == '1' || bit == '0' {
					bits[i] = mask[i]
				}
			}
			mem[addr] = string(bits)
		}
	}

	return fmt.Sprintf("%d", sumMem(mem))
}

// Day14b solves the puzzle and returns a string
func Day14b(input []string) string {

	mem := make(map[int]string)
	mask := strings.Repeat("X", 36)
	for _, instruction := range input {
		parts := strings.Split(instruction, " = ")
		if parts[0] == "mask" {
			mask = parts[1]
		} else {
			valBits := binStringFromDec(parts[1])
			addrBinStr := []byte(binStringFromDec(memRegex.FindString(parts[0])))

			var replacementIndicies []int

			for i, bit := range mask {
				if bit == '1' {
					addrBinStr[i] = '1'
				} else if bit == 'X' {
					replacementIndicies = append(replacementIndicies, i)
				}
			}

			if len(replacementIndicies) > 0 {
				max := 1 << len(replacementIndicies)
				for i := 0; i < max; i++ {
					tmpFmt := fmt.Sprint("%0", len(replacementIndicies), "b")
					tmp := fmt.Sprintf(tmpFmt, int64(i))

					for j := 0; j < len(replacementIndicies); j++ {
						addrBinStr[replacementIndicies[j]] = tmp[j]
					}
					tmpNumAddr, _ := strconv.ParseInt(string(addrBinStr), 2, 64)
					mem[int(tmpNumAddr)] = valBits
				}
			} else {
				tmpNumAddr, _ := strconv.ParseInt(string(addrBinStr), 2, 64)
				mem[int(tmpNumAddr)] = valBits
			}
		}
	}

	return fmt.Sprintf("%d", sumMem(mem))
}
