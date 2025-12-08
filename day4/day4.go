package day4

import (
	"aoc2025/utils"
	"fmt"
	"log"
	"strings"
)

func Day4() {
	part1()
	part2()
}

func checkSurrounding(matrix [][]string, i int, j int) int {
	count := 0
	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			if k == i && l == j {
				continue
			}
			if k < 0 || k >= len(matrix) || l < 0 || l >= len(matrix[k]) {
				continue
			}
			if matrix[k][l] == "@" {
				count++
			}
		}
	}
	return count
}

func part1() {
	lines, err := utils.ReadInputFile("day4/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	paperCount := 0
	matrix := make([][]string, 0)
	for _, line := range lines {
		split := strings.Split(line, "")
		matrix = append(matrix, split)
	}

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == "@" {
				if checkSurrounding(matrix, i, j) < 4 {
					paperCount++
				}
			}
		}
	}

	fmt.Printf("Part 1 Rolls Accessed: %d\n", paperCount)
}

func part2() {
	lines, err := utils.ReadInputFile("day4/day4.txt")
	if err != nil {
		log.Fatal(err)
	}
	paperCount := 0
	matrix := make([][]string, 0)
	for _, line := range lines {
		split := strings.Split(line, "")
		matrix = append(matrix, split)
	}
	change := true
	for change {
		iterCount := 0
		for i := 0; i < len(matrix); i++ {
			for j := 0; j < len(matrix[i]); j++ {
				if matrix[i][j] == "@" {
					if checkSurrounding(matrix, i, j) < 4 {
						iterCount++
						paperCount++
						matrix[i][j] = "x"
					}
				}
			}
		}
		if iterCount == 0 {
			change = false
		}
	}

	fmt.Printf("Part 1 Rolls Accessed: %d\n", paperCount)
}
