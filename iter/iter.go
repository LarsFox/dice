package iter

import "math/rand"

// Product returns a cartesian product of k from 0 to n
func Product(n, k int) [][]int {
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

// ShuffleIntInt randomly permutates a slice of slices of int
func ShuffleIntInt(slice [][]int) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}

// ShuffleInt randomly permutates a slice of int
func ShuffleInt(slice []int) {
	for i := range slice {
		j := rand.Intn(i + 1)
		slice[i], slice[j] = slice[j], slice[i]
	}
}
