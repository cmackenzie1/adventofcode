package days

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"github.com/cmackenzie1/adventofcode/v2022/input"
)

func Day1Part1(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	elves := make(map[int]int)
	i := 1
	for _, line := range lines {
		if line == "" {
			i += 1
			continue
		}

		n, err := strconv.Atoi(line)
		if err != nil {
			log.Fatalf("unable to process line: %q", line)
		}

		if _, ok := elves[i]; !ok {
			elves[i] = n
		} else {
			elves[i] += n
		}

	}

	var maxNum int
	var maxElf int
	for k, v := range elves {
		if v > maxNum {
			maxElf = k
			maxNum = v
		}
	}

	fmt.Printf("elf #%d has %d calories\n", maxElf, maxNum)
}

type Elf struct {
	id            int
	snacks        []int
	totalCalories int
}

func (e *Elf) String() string {
	return fmt.Sprintf("Elf{id: %d, snacks: %v, totalCalories: %d}", e.id, e.snacks, e.totalCalories)
}

type elfSorter struct {
	elves []*Elf
	by    func(e1, e2 *Elf) bool
}

type By func(e1, e2 *Elf) bool

func (by By) Sort(elves []*Elf) {
	es := &elfSorter{
		elves: elves,
		by:    by,
	}
	sort.Sort(es)
}

func (es *elfSorter) Len() int {
	return len(es.elves)
}

func (es *elfSorter) Less(i, j int) bool {
	return es.by(es.elves[i], es.elves[j])
}

func (es *elfSorter) Swap(i, j int) {
	es.elves[i], es.elves[j] = es.elves[j], es.elves[i]
}

func Day1Part2(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	elves := make([]*Elf, 0, 0)
	id := 1
	snacks := make([]int, 0)
	for _, line := range lines {
		if line == "" {
			elf := &Elf{id: id, snacks: snacks[:], totalCalories: SumInts(snacks)}
			elves = append(elves, elf)
			log.Print(elf)
			id += 1
			snacks = make([]int, 0)
			continue
		}

		n, _ := strconv.Atoi(line)
		snacks = append(snacks, n)
	}

	var calories = func(e1, e2 *Elf) bool {
		return e1.totalCalories < e2.totalCalories
	}

	By(calories).Sort(elves)
	fmt.Printf("top 3 elves: %s\n", elves[:3])
	var sum int
	for i := 0; i < 3; i++ {
		sum += elves[len(elves)-i-1].totalCalories
	}
	fmt.Printf("top 3 elves sum: %d\n", sum)

}

func SumInts(input []int) int {
	var s int
	for _, v := range input {
		s += v
	}
	return s
}
