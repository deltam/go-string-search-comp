package main

import (
	"fmt"
)

func main() {
	s := "ABC ABCDAB ABCDABCDABD"
	sub := "ABCDABD"

	counter = 0
	fmt.Println("RabinKarp", indexRabinKarp(s, sub), "comp", counter)
	counter = 0
	fmt.Println("Naive    ", indexNaive(s, sub), "comp", counter)
	counter = 0
	fmt.Println("KMP      ", indexKMP(s, sub), "comp", counter)
}

var counter int

// primeRK is the prime base used in Rabin-Karp algorithm.
const primeRK = 16777619

// hashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashStr(sep string) (uint32, uint32) {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	var pow, sq uint32 = 1, primeRK
	for i := len(sep); i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return hash, pow
}

// copy form pkg/strings
func indexRabinKarp(s, substr string) int {
	// Rabin-Karp search
	hashss, pow := hashStr(substr)
	n := len(substr)
	var h uint32
	for i := 0; i < n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if h == hashss && s[:n] == substr {
		return 0
	}
	for i := n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= pow * uint32(s[i-n])
		i++
		counter++
		if h == hashss && s[i-n:i] == substr {
			return i - n
		}
	}
	return -1
}

// Naive
func indexNaive(s string, substr string) int {
	i, j := 0, 0
	for i < len(s) && j < len(substr) {
		if s[i] != substr[j] {
			i -= j - 1
			j = 0
			counter++
			continue
		}
		i++
		j++
	}
	if j == len(substr) {
		return i - j
	}
	return -1
}

// Knuth-Morris-Pratt
func indexKMP(s, substr string) int {
	next := make([]int, len(substr))
	next[0] = -1
	i, j := 0, -1
	for {
		for j >= 0 && substr[i] != substr[j] {
			j = next[j]
		}
		i++
		j++
		if i >= len(substr) {
			break
		}
		if substr[i] == substr[j] {
			next[i] = next[j]
		} else {
			next[i] = j
		}
	}

	i, j = 0, 0
	for i < len(s) && j < len(substr) {
		for j >= 0 && s[i] != substr[j] {
			j = next[j]
			counter++
		}
		i++
		j++
	}
	if j == len(substr) {
		return i - j
	}
	return -1
}

