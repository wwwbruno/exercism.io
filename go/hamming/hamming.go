package hamming

import (
	"errors"
	"strings"
)

const testVersion = 6

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("Different sizes")
	}

	var count int
	dna1 := strings.Split(a, "")
	dna2 := strings.Split(b, "")

	for i := range dna1 {
		if dna1[i] != dna2[i] {
			count++
		}
	}

	return count, nil
}
