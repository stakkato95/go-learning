package main

import (
	"fmt"
	"testing"
)

func BenchmarkFewLoops(b *testing.B) {
	for i := 0; i < 10; i++ {
		fmt.Sprintln(i)
	}
}

func BenchmarkManyLoops(b *testing.B) {
	for i := 0; i < 1_000_000; i++ {
		fmt.Sprintln(i)
	}
}
