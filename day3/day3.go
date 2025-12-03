package day3

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"slices"
	"strconv"
	"strings"
)

func Day3() {
	part1()
	part2()
}

func part1() {
	lines, err := utils.ReadInputFile("day3/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	totalVoltage := 0
	for _, line := range lines {
		length := len(line)
		split := strings.Split(line, "")
		numSlice := make([]int, 0)
		for _, val := range split {
			numeric, _ := strconv.Atoi(val)
			numSlice = append(numSlice, numeric)
		}
		largestIndex := 0
		for i := 0; i < length-1; i++ {
			if numSlice[i] > numSlice[largestIndex] {
				largestIndex = i
			}
		}
		nextLargeIndex := largestIndex + 1
		for j := nextLargeIndex; j < length; j++ {
			if numSlice[j] > numSlice[nextLargeIndex] {
				nextLargeIndex = j
			}
		}
		joined := split[largestIndex] + split[nextLargeIndex]
		voltage, _ := strconv.Atoi(joined)
		totalVoltage += voltage
	}
	fmt.Printf("Part 1 Max Voltage: %d\n", totalVoltage)
}

func part2() {
	lines, err := utils.ReadInputFile("day3/day3.txt")
	if err != nil {
		log.Fatal(err)
	}
	totalVoltage := 0
	for _, line := range lines {
		length := len(line)
		split := strings.Split(line, "")
		numSlice := make([]int, 0)
		for _, val := range split {
			numeric, _ := strconv.Atoi(val)
			numSlice = append(numSlice, numeric)
		}
		largestIndex := 0
		for i := 0; i < length-11; i++ {
			if numSlice[i] > numSlice[largestIndex] {
				largestIndex = i
			}
		}

		storedIndexes := make([]int, 12)
		storedIndexes[0] = largestIndex
		counter := 1

		for counter < 12 {
			found := 0
			until := length - 11 + counter
			for i := storedIndexes[counter-1] + 1; i < until; i++ {
				if numSlice[i] > found && i > storedIndexes[counter-1] {
					storedIndexes[counter] = i
					found = numSlice[i]
				}
			}
			counter++
		}
		slices.Sort(storedIndexes)
		var numberString string
		for _, key := range storedIndexes {
			numberString += split[key]
		}
		strNum, _ := strconv.Atoi(numberString)
		totalVoltage += strNum
	}
	fmt.Printf("Part 2 Max Voltage: %d\n", totalVoltage)
}
