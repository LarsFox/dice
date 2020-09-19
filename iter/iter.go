package iter

import (
	"math/rand"
	"time"
)

// Sequence generates a new sequence of iterations, where:
// n is a number of values for one die, usually 6.
// k is a number of dice, usually 1–3.
// s is a number of repetitions of dice sequences, 3 is usually enough.
func Sequence(n, k, s int) [][]int {
	rand.Seed(time.Now().UnixNano())
	seq := make([][]int, 0, n^k*s)
	for i := 0; i < s; i++ {
		for _, perm := range product(n, k) {
			shuffleInt(perm)
			seq = append(seq, perm)
		}
	}
	shuffleIntInt(seq)
	return seq
}

// product returns a cartesian product of k from 0 to n.
func product(n, k int) [][]int {
	result := make([][]int, 0, n^k)
	ix := make([]int, k)
	for {
		pair := make([]int, k)
		copy(pair, ix)
		result = append(result, pair)
		j := k - 1
		for ; j >= 0 && ix[j] == n-1; j-- {
			ix[j] = 0
		}
		if j < 0 {
			return result
		}
		ix[j]++
	}
}

// shuffleIntInt randomly permutates a slice of slices of int.
func shuffleIntInt(slice [][]int) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// shuffleInt randomly permutates a slice of int.
func shuffleInt(slice []int) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
