package containers

import "fmt"

type Counter[K comparable] struct {
	m map[K]int
}

func (s *Counter[K]) String() string {
	return fmt.Sprintf("%v", s.m)
}

func NewCounter[K comparable]() *Counter[K] {
	return &Counter[K]{m: make(map[K]int)}
}

func (s *Counter[K]) Add(k K) {
	if _, ok := s.m[k]; ok {
		s.m[k] += 1
	} else {
		s.m[k] = 1
	}
}

func (s *Counter[K]) AddAll(v ...K) {
	for _, k := range v {
		s.Add(k)
	}
}

func (s *Counter[K]) Remove(k K) {
	delete(s.m, k)
}

func (s *Counter[K]) Contains(k K) bool {
	_, c := s.m[k]
	return c
}

func (s *Counter[K]) Total(k K) int {
	return s.m[k]
}
