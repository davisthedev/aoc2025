package day2

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Day2() {
	part1()
	part2()
}

func isInvalid(num int) bool {
	length := len(strconv.Itoa(num))
	if length%2 == 0 {
		middle := length / 2
		asStr := strconv.Itoa(num)
		return asStr[:middle] == asStr[middle:]
	}
	return false
}

func part1() {
	lines, err := utils.ReadInputCsvFile("day2/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	var total int
	for _, line := range lines {
		split := strings.Split(line, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		for i := start; i <= end; i++ {
			if isInvalid(i) {
				total += i
			}
		}
	}
	fmt.Printf("Part 1 Invalid Sum: %d\n", total)
}

func isMultiInvalid(num int) bool {
	numStr := strconv.Itoa(num)
	length := len(numStr)
	if length > 1 {
		for i := 2; i <= length; i++ {
			chunkSize := length / i
			pattern := numStr[0:chunkSize]
			expected := strings.Repeat(pattern, i)
			if expected == numStr {
				return true
			}
		}
	}
	return false
}

func part2() {
	lines, err := utils.ReadInputCsvFile("day2/day2.txt")
	if err != nil {
		log.Fatal(err)
	}
	var total int
	for _, line := range lines {
		split := strings.Split(line, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		for i := start; i <= end; i++ {
			if isMultiInvalid(i) {
				total += i
			}
		}
	}
	fmt.Printf("Part 2 Invalid Sum: %d\n", total)
}
