package days

import "testing"

func TestDay3_Rucksack(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "vJrwpWtwJgWrhcsFMMfFFhFp", want: "p"},
		{"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL", "L"},
		{"PmmdzqPrVvPwwTWBwg", "P"},
		{"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn", "v"},
		{"ttgJtRGJQctTZtZT", "t"},
		{"CrZsJsPPZsGzwwsLwLmpwMDw", "s"},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			r := NewRucksack(tt.input)
			if got := r.Common(); got != tt.want {
				t.Errorf("Common() want = %s, got = %s", tt.want, got)
			}
		})
	}
}

func TestDay3_Weight(t *testing.T) {
	tests := []struct {
		r    rune
		want int
	}{
		{r: 'a', want: 1},
		{r: 'z', want: 26},
		{r: 'A', want: 27},
		{r: 'Z', want: 52},
	}

	for _, tt := range tests {
		t.Run(string(tt.r), func(t *testing.T) {
			if got := weight(string(tt.r)); got != tt.want {
				t.Errorf("weight() want = %d, got = %d", tt.want, got)
			}
		})
	}
}
