package goBenchExp

import (
	"fmt"
	"math"
	"testing"
)

func BenchmarkInlineFunc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		inlineFunc(1, 2)
	}
}

func BenchmarkEmptyFuncVs0(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if i == math.MaxInt32 {
			fmt.Println(0)
		}
	}
}

func BenchmarkEmptyFunc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		emptyFunc(0)
	}
}

func BenchmarkGoEmptyFunc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		go emptyFunc(0)
	}
}

func BenchmarkSimpleDefer(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		simpleDefer(0)
	}
}
