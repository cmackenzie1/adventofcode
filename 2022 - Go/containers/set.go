package containers

import (
	"fmt"
)

type Set[K comparable] struct {
	m map[K]struct{}
}

func (s *Set[K]) String() string {
	return fmt.Sprintf("%v", s.m)
}

func NewSet[K comparable]() *Set[K] {
	return &Set[K]{m: make(map[K]struct{})}
}

// Add inserts an element into the set if it does not already exist
func (s *Set[K]) Add(k K) {
	s.add(k)
}

func (s *Set[K]) add(k K) {
	s.m[k] = struct{}{}
}

// AddAll inserts each element of a "splat-able" slice into the set.
// A "splat-able" slice is a slice where the `...` operator does not modify the type.
func (s *Set[K]) AddAll(v ...K) {
	for _, k := range v {
		s.add(k)
	}
}

// Remove deletes an element from the set
func (s *Set[K]) Remove(k K) {
	delete(s.m, k)
}

// Values returns a slice of all the elements in the set
func (s *Set[K]) Values() []K {
	values := make([]K, 0)
	for v := range s.m {
		values = append(values, v)
	}

	return values
}

// Intersects returns a new set representing the intersection
// of two sets.
func (s *Set[K]) Intersects(k *Set[K]) *Set[K] {
	intersection := NewSet[K]()
	for _, v := range s.Values() {
		if k.Contains(v) {
			intersection.Add(v)
		}
	}
	return intersection
}

// Contains returns true if the element is present in the set.
// Otherwise, false.
func (s *Set[K]) Contains(k K) bool {
	_, c := s.m[k]
	return c
}
