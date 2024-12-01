package days

import (
	"log"

	"github.com/cmackenzie1/adventofcode/v2022/containers"
	"github.com/cmackenzie1/adventofcode/v2022/input"
)

func Day6Part1(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	stream := lines[0]
	windowSize := 4
	for i := 0; i < len(stream)-windowSize; i++ {
		chunk := stream[i : i+windowSize]
		if unique(chunk) {
			log.Printf("sequence = %s, chars = %d", chunk, i+windowSize)
			break
		}
	}
}

func unique(s string) bool {
	seen := containers.NewSet[rune]()
	for _, r := range s {
		if seen.Contains(r) {
			return false
		}
		seen.Add(r)
	}
	return true
}

func Day6Part2(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	stream := lines[0]
	windowSize := 14
	for i := 0; i < len(stream)-windowSize; i++ {
		chunk := stream[i : i+windowSize]
		if unique(chunk) {
			log.Printf("sequence = %s, chars = %d", chunk, i+windowSize)
			break
		}
	}
}
