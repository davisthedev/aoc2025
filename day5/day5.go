package day5

import (
	"aoc2025/utils"
	"fmt"
	"log"
)

func Day5() {
	part1()
	part2()
}

func part1() {
	lines, err := utils.ReadInputFile("day4/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	something := 0
	for _, line := range lines {
		log.Println(line)
	}

	fmt.Printf("Part 1 Rolls Accessed: %d\n", something)
}

func part2() {

}
