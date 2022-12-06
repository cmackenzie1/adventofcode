package days

import (
	"fmt"
	"log"

	"github.com/cmackenzie1/adventofcode/v2022/input"
)

func Day5Part1(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	// split the input into two chunks of lines, the stacks and
	// the movement instructions
	stacksInput := make([]string, 0)
	instructionsInput := make([]string, 0)
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			stacksInput = append(stacksInput, lines[:i]...)
			instructionsInput = append(instructionsInput, lines[i+1:]...) // i+1 to drop the empty line
			break
		}
	}

	log.Printf("total crates: %#v", stacksInput)
	log.Printf("total instructions: %#v", instructionsInput)

	cols := make([][]string, 0)
	for _, line := range stacksInput {
		cols = append(cols, crates(line))
	}

	// stacks is a slice of slices. Each element of the first slice represents a stack
	// and each element in the stack are the crates themselves. It's important to note
	// the stack of crates behaves like the Stack data structure, where only the last element is
	// poppable.
	stacks := make([][]string, len(cols))
	for i := 0; i < len(cols)-1; i++ {
		for j := 0; j < len(cols[i]); j++ {
			if cols[i][j] != " " { // drop blanks
				stacks[j] = append([]string{cols[i][j]}, stacks[j]...)
			}
		}
	}

	log.Printf("crates: %#v", stacks)
	for _, line := range instructionsInput {
		var delta, from, to int
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &delta, &from, &to)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("delta = %d, from = %d, to = %d", delta, from, to)

		crateMover9000(delta, from-1, to-1, stacks)
	}

	log.Printf("final stacks: %#v", stacks)
	var s string
	for _, stack := range stacks {
		s += stack[len(stack)-1]
	}
	log.Printf("final stacks: %s", s)
}

func crates(s string) []string {
	result := make([]string, 0)
	// crates are at positions: 1, 5, 9, 13, 17, 21
	for i := 1; i < len(s); i += 4 {
		result = append(result, string(s[i]))
	}
	return result
}

func Day5Part2(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	// split the input into two chunks of lines, the stacks and
	// the movement instructions
	stacksInput := make([]string, 0)
	instructionsInput := make([]string, 0)
	for i := 0; i < len(lines); i++ {
		if lines[i] == "" {
			stacksInput = append(stacksInput, lines[:i]...)
			instructionsInput = append(instructionsInput, lines[i+1:]...) // i+1 to drop the empty line
			break
		}
	}

	log.Printf("total crates: %#v", stacksInput)
	log.Printf("total instructions: %#v", instructionsInput)

	cols := make([][]string, 0)
	for _, line := range stacksInput {
		cols = append(cols, crates(line))
	}

	// stacks is a slice of slices. Each element of the first slice represents a stack
	// and each element in the stack are the crates themselves. It's important to note
	// the stack of crates behaves like the Stack data structure, where only the last element is
	// poppable.
	stacks := make([][]string, len(cols))
	for i := 0; i < len(cols)-1; i++ {
		for j := 0; j < len(cols[i]); j++ {
			if cols[i][j] != " " { // drop blanks
				stacks[j] = append([]string{cols[i][j]}, stacks[j]...)
			}
		}
	}

	log.Printf("crates: %#v", stacks)
	for _, line := range instructionsInput {
		var delta, from, to int
		_, err := fmt.Sscanf(line, "move %d from %d to %d", &delta, &from, &to)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("================================================================================")
		log.Printf("before = %#v", stacks)
		log.Printf("delta = %d, from = %d, to = %d", delta, from, to)
		crateMover9001(delta, from-1, to-1, stacks)
		log.Printf("after = %#v", stacks)

	}

	log.Printf("final stacks: %#v", stacks)
	var s string
	for _, stack := range stacks {
		s += stack[len(stack)-1]
	}
	log.Printf("final stacks: %s", s)
}

func crateMover9000(n int, from int, to int, stacks [][]string) {
	for i := 0; i < n; i++ {
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
		stacks[from] = stacks[from][:len(stacks[from])-1]
	}
}

func crateMover9001(n int, from int, to int, stacks [][]string) {
	fromIndex := len(stacks[from]) - n
	endIndex := len(stacks[from])
	pick := stacks[from][fromIndex:endIndex]
	log.Printf("from = %#v, take = %#v", stacks[from], pick)
	stacks[to] = append(stacks[to], stacks[from][fromIndex:endIndex]...)
	stacks[from] = stacks[from][:fromIndex]
}
