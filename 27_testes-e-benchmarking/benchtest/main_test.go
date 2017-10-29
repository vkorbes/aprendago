package main

import "testing"

// func TestAdd (t *testing.T)

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Add(1, 2)
	}
}
