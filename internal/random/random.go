// Package random provides functions for generating random numbers.
package random

import (
	"math/rand"
)

// Range generates a random integer within the specified range [minimum, maximum].
func Range(minimum, maximum int) int {
	length := maximum - minimum
	if length <= 0 {
		panic("invalid argument to Range")
	}
	vector := generateVector(length, minimum)
	return vector[rand.Intn(length)]
}

// generateVector generates a vector of integers starting from the initial value.
func generateVector(length, initialValue int) []int {
	vector := make([]int, length)

	for i := 0; i < length; i++ {
		vector[i] = initialValue + i
	}
	return vector
}
