package main

import (
	"math/rand"
	"testing"
)

func makeBenchInputHard() string {
	tokens := [...]string{
		"<a>", "<p>", "<b>", "<strong>",
		"</a>", "</p>", "</b>", "</strong>",
		"hello", "world",
	}
	x := make([]byte, 0, 1<<20)
	for {
		i := rand.Intn(len(tokens))
		if len(x)+len(tokens[i]) >= 1<<20 {
			break
		}
		x = append(x, tokens[i]...)
	}
	return string(x)
}

var benchInputHard = makeBenchInputHard()

func benchmarkIndexHard(b *testing.B, f func(string, string) int, sep string) {
	for i := 0; i < b.N; i++ {
		f(benchInputHard, sep)
	}
}

func BenchmarkRabinKarp(b *testing.B) {
	benchmarkIndexHard(b, indexRabinKarp, "<pre><b>hello</b><strong>world</strong></pre>")
}

func BenchmarkNaive(b *testing.B) {
	benchmarkIndexHard(b, indexNaive, "<pre><b>hello</b><strong>world</strong></pre>")
}

func BenchmarkKMP(b *testing.B) {
	benchmarkIndexHard(b, indexKMP, "<pre><b>hello</b><strong>world</strong></pre>")
}

