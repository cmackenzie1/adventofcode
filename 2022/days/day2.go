package days

import (
	"log"
	"strings"

	"github.com/cmackenzie1/adventofcode/v2022/input"
)

type Round struct {
	left  string
	right string
}

func Day2Part1(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	rounds := make([]Round, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}

		splits := strings.Split(line, " ")
		rounds = append(rounds, Round{left: splits[0], right: splits[1]})
	}

	var totalScore int
	for i, r := range rounds {
		winner, score := Play(r.left, r.right)
		totalScore = totalScore + score

		log.Printf("round %d: left = %s, right = %s, winner = %s, score = %d\n", i, r.left, r.right, winner, score)
	}

	log.Printf("total score = %d\n", totalScore)
}

const (
	winScore  = 6
	drawScore = 3
	loseScore = 0

	rockScore     = 1
	paperScore    = 2
	scissorsScore = 3
)

// Play determines the winner and resulting score for a round
// of rock, paper, scissors.
// A -> Rock <- X
// B -> Paper <- Y
// C -> Scissors <- Z
func Play(l, r string) (result string, score int) {
	switch l {
	case "A": // rock
		switch r {
		case "X": // rock == draw
			return "draw", rockScore + drawScore
		case "Y": // paper == win
			return "right", paperScore + winScore
		case "Z": // scissors == lose
			return "left", scissorsScore + loseScore
		}
	case "B": // paper
		switch r {
		case "X": // rock == lose
			return "left", rockScore + loseScore
		case "Y": // paper == draw
			return "draw", paperScore + drawScore
		case "Z": // scissors == win
			return "right", scissorsScore + winScore
		}
	case "C": // scissors
		switch r {
		case "X": // rock == win
			return "right", rockScore + winScore
		case "Y": // paper == lose
			return "left", paperScore + loseScore
		case "Z": // scissors == draw
			return "draw", scissorsScore + drawScore
		}
	}

	return "unknown", 0
}

const (
	rock     = "X"
	paper    = "Y"
	scissors = "Z"
)

func (r *Round) Pick() string {
	switch r.right {
	case "X": // need to lose
		switch r.left {
		case "A":
			return scissors
		case "B":
			return rock
		case "C":
			return paper
		}
	case "Y": // need to draw
		switch r.left {
		case "A":
			return rock
		case "B":
			return paper
		case "C":
			return scissors
		}
	case "Z": // need to win
		switch r.left {
		case "A":
			return paper
		case "B":
			return scissors
		case "C":
			return rock
		}
	}
	return "unknown"
}

func Day2Part2(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	rounds := make([]Round, 0)
	for _, line := range lines {
		if line == "" {
			continue
		}

		splits := strings.Split(line, " ")
		rounds = append(rounds, Round{left: splits[0], right: splits[1]})
	}

	var totalScore int
	for i, r := range rounds {
		winner, score := Play(r.left, r.Pick())
		totalScore = totalScore + score

		log.Printf("round %d: left = %s, right = %s, winner = %s, score = %d\n", i, r.left, r.right, winner, score)
	}

	log.Printf("total score = %d\n", totalScore)
}
