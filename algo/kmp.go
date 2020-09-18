package algo

import (
	log "github.com/sirupsen/logrus"
)

func computePrefix(pattern string) []int {
	m := len(pattern)
	pi := make([]int, m)
	k := -1

	pi[0] = 0
	for q := 1; q < m; q++ {
		for k >= 0 && pattern[k+1] != pattern[q] {
			k = pi[k] - 1 // pi contains length of prefix
		}

		if pattern[k+1] == pattern[q] {
			k++
		}

		pi[q] = k + 1 // pi contains length of prefix
	}

	return pi
}

func displayPi(pattern string, prefixArray []int) {
	for _, val := range pattern {
		log.Debug(val)
	}
	log.Debug("\n")

	for _, val := range prefixArray {
		log.Debug(val)
	}
	log.Debug("\n")
}

// KmpMatcher .. prints the occurance of 'pattern' in 'text'
// Implementation of KMP algorithm
func KmpMatcher(text, pattern string) []int {
	result := []int{}
	n := len(text)
	m := len(pattern)

	pi := computePrefix(pattern)
	displayPi(pattern, pi)

	k := -1 // number of characters matched

	for i := 0; i < n; i++ { //scan text from left to right
		for k >= 0 && pattern[k+1] != text[i] {
			k = pi[k] - 1
		}

		if pattern[k+1] == text[i] {
			k++
		}

		if k+1 == m {
			log.Debug("Pattern matched with shift ", i-m+1)
			result = append(result, i-m+1)
			k = pi[k] - 1
		}
	}

	return result
}
