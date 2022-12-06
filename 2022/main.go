package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"

	"github.com/cmackenzie1/adventofcode/v2022/days"
)

var solutions = map[string]func(input string){
	"d1p1": days.Day1Part1,
	"d1p2": days.Day1Part2,

	"d2p1": days.Day2Part1,
	"d2p2": days.Day2Part2,

	"d3p1": days.Day3Part1,
	"d3p2": days.Day3Part2,

	"d4p1": days.Day4Part1,
	"d4p2": days.Day4Part2,

	"d5p1": days.Day5Part1,
	"d5p2": days.Day5Part2,
}

func main() {
	day := flag.Int("day", 0, "What day to run")
	part := flag.Int("part", 0, "What part to run")
	flag.Parse()

	dp := fmt.Sprintf("d%dp%d", *day, *part)
	if _, ok := solutions[dp]; !ok {
		log.Fatalf("Invalid selection!")
	}

	path, err := filepath.Abs(filepath.Join("input", fmt.Sprintf("day%d.txt", *day)))
	if err != nil {
		log.Fatal(err)
	}

	solutions[dp](path)
}
