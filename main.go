package main

import (
	"aoc2025/day1"
	"aoc2025/day2"
	"aoc2025/day3"
	"aoc2025/day4"
	"aoc2025/day5"
	"flag"
	"fmt"
	"os"
	"time"
)

func main() {
	start := time.Now()
	day := flag.Int("day", 0, "Advent of Code day to run (1-12)")
	flag.Parse()

	if *day < 1 || *day > 12 {
		fmt.Println("Error: Enter a day between 1 and 12")
		flag.Usage()
		os.Exit(1)
	}

	days := map[int]func(){
		1: day1.Day1,
		2: day2.Day2,
		3: day3.Day3,
		4: day4.Day4,
		5: day5.Day5,
	}

	if fn, exists := days[*day]; exists {
		fmt.Printf("Running Day %d...\n\n", *day)
		fn()

	} else {
		fmt.Printf("Day %d not implemented\n", *day)
		os.Exit(1)
	}

	elapsed := time.Since(start)
	fmt.Printf("\nCompleted in %s\n", elapsed)
}
