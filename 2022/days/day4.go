package days

import (
	"log"
	"strconv"
	"strings"

	"github.com/cmackenzie1/adventofcode/v2022/input"
)

type InclusiveRange struct {
	low  int
	high int
}

func (r *InclusiveRange) Contains(i int) bool {
	return i >= r.low && i <= r.high
}

func Day4Part1(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	totalOverlaps := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		splits := strings.Split(line, ",")
		left, right := NewInclusiveRangeFromString(splits[0]), NewInclusiveRangeFromString(splits[1])

		if (left.Contains(right.low) && left.Contains(right.high)) || (right.Contains(left.low) && right.Contains(left.high)) {
			totalOverlaps += 1
		}
	}

	log.Printf("total overlaps: %d", totalOverlaps)
}

// NewInclusiveRangeFromString takes a string consisting of a range
// separated by a dash. For example, the range 2 through 6 can be
// represented as 2-6
func NewInclusiveRangeFromString(r string) *InclusiveRange {
	splits := strings.Split(r, "-")
	low, _ := strconv.Atoi(splits[0])
	high, _ := strconv.Atoi(splits[1])
	return &InclusiveRange{
		low:  low,
		high: high,
	}
}

func Day4Part2(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	totalOverlaps := 0
	for _, line := range lines {
		if line == "" {
			continue
		}
		splits := strings.Split(line, ",")
		left, right := NewInclusiveRangeFromString(splits[0]), NewInclusiveRangeFromString(splits[1])

		if (left.Contains(right.low) || left.Contains(right.high)) || (right.Contains(left.low) || right.Contains(left.high)) {
			totalOverlaps += 1
		}
	}

	log.Printf("total overlaps: %d", totalOverlaps)
}
