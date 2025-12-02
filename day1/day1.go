package day1

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"math"
	"strconv"
)

func Day1() {
	part1()
	part2()
}

func part1() {
	lines, err := utils.ReadInputFile("day1/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	dial := 50
	zeroCount := 0
	for _, line := range lines {
		direction := line[0]
		amount, _ := strconv.Atoi(line[1:])

		switch direction {
		case 'L':
			dial = ((dial-amount)%100 + 100) % 100
		case 'R':
			dial = (dial + amount) % 100
		}
		if dial == 0 {
			zeroCount++
		}
	}
	fmt.Printf("Part 1 Zero Count: %d\n", zeroCount)
}

func part2() {
	lines, err := utils.ReadInputFile("day1/day1.txt")
	if err != nil {
		log.Fatal(err)
	}
	dial := 50
	zeroCount := 0
	for _, line := range lines {
		direction := line[0]
		amount, _ := strconv.Atoi(line[1:])

		var change int
		switch direction {
		case 'L':
			change = -amount
		case 'R':
			change = amount
		}

		var previousRotation float64
		if change < 0 {
			previousRotation = math.Ceil(float64(dial) / 100.00)
		} else {
			previousRotation = math.Floor(float64(dial) / 100.00)
		}

		dial += change

		var currentRotation float64
		if change < 0 {
			currentRotation = math.Ceil(float64(dial) / 100.00)
		} else {
			currentRotation = math.Floor(float64(dial) / 100.00)
		}

		crossings := int(math.Abs(currentRotation - previousRotation))
		zeroCount += crossings
	}
	fmt.Printf("Part 2 Zero Count: %d\n", zeroCount)
}
