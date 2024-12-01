package days

import (
	"log"
	"strings"

	"github.com/cmackenzie1/adventofcode/v2022/containers"
	"github.com/cmackenzie1/adventofcode/v2022/input"
)

type Rucksack struct {
	items string
	c1    string
	c2    string
}

func NewRucksack(items string) *Rucksack {
	mid := len(items) / 2
	c1, c2 := items[:mid], items[mid:]
	return &Rucksack{items: items, c1: c1, c2: c2}
}

func (r *Rucksack) Common() string {
	s := containers.NewSet[int32]()
	for _, v := range r.c1 {
		s.Add(v)
	}

	for _, v := range r.c2 {
		if s.Contains(v) {
			return string(v)
		}
	}

	return ""
}

func (r *Rucksack) Items() string {
	return r.items
}

func (r *Rucksack) Set() *containers.Set[rune] {
	s := containers.NewSet[rune]()
	for _, v := range r.Items() {
		s.Add(v)
	}
	return s
}

const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func weight(l string) int {
	return strings.Index(alphabet, l) + 1
}

func Day3Part1(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	rucksacks := make([]*Rucksack, 0)
	for _, line := range lines {
		rucksacks = append(rucksacks, NewRucksack(line))
	}

	log.Printf("total rucksacks: %d", len(rucksacks))
	var total int
	for _, r := range rucksacks {
		c := r.Common()
		w := weight(c)
		log.Printf("common items = %v, score = %d", c, w)
		total += w
	}

	log.Printf("total score: %d", total)
}

func Day3Part2(path string) {
	lines, err := input.ReadLines(path)
	if err != nil {
		log.Fatal(err)
	}

	rucksacks := make([]*Rucksack, 0)
	for _, line := range lines {
		rucksacks = append(rucksacks, NewRucksack(line))
	}

	groups := make([][]*Rucksack, 0)
	group := make([]*Rucksack, 0, 3)
	for _, r := range rucksacks {
		group = append(group, r)
		if len(group) == 3 {
			groups = append(groups, group)
			group = make([]*Rucksack, 0, 3)
		}
	}

	badges := make([]rune, 0)
	for _, gr := range groups {
		intersection := gr[0].Set()
		for i := 1; i < len(gr); i++ {
			intersection = intersection.Intersects(gr[i].Set())
		}

		log.Printf("intersection: %v", string(intersection.Values()[0]))
		badges = append(badges, intersection.Values()[0])
	}

	var total int
	for _, v := range badges {
		total += weight(string(v))
	}
	log.Printf("total : %d", total)
}
