package containers

import (
	"testing"
)

func TestCounter_Add(t *testing.T) {
	tests := map[string]struct {
		name    string
		values  []string
		checkFn func(c *Counter[string]) bool
	}{
		"single input": {values: []string{"a"}, checkFn: func(c *Counter[string]) bool {
			return c.m["a"] == 1
		}},
		"count same": {values: []string{"a", "a", "a"}, checkFn: func(c *Counter[string]) bool {
			return c.m["a"] == 3
		}},
		"count distinct": {values: []string{"a", "b", "c"}, checkFn: func(c *Counter[string]) bool {
			return c.m["a"] == 1 && c.m["b"] == 1 && c.m["c"] == 1
		}},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			c := NewCounter[string]()
			for _, v := range tt.values {
				c.Add(v)
			}

			if !tt.checkFn(c) {
				t.Errorf("Add() did not properly increment values")
			}
		})
	}
}
