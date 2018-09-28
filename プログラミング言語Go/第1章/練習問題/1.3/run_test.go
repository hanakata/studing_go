package main

import (
	"strings"
	"testing"
)

func Benchmark_Fast(b *testing.B) {
	b.ResetTimer()
	s := []string{"aaa", "bbb", "ccc", "ddd"}
	var sep, t, a string
	sep = " "
	for i := 0; i < len(s); i++ {
		t = s[i]
		a += sep + t
	}
}

func Benchmark_Delay(b *testing.B) {
	b.ResetTimer()
	s := []string{"aaa", "bbb", "ccc", "ddd"}
	strings.Join(s[0:], " ")
}
