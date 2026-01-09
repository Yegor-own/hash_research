package rk

import (
	"time"

	"rabin_karp_exp/hash"
)

type Metrics struct {
	Time            time.Duration
	HashMatches     int
	Collisions      int
	CharComparisons int
	Matches         int
}

func Search(text, pattern string, h hash.Hasher) Metrics {
	n, m := len(text), len(pattern)
	var metrics Metrics

	if m > n {
		return metrics
	}

	start := time.Now()

	patternHash := h.Init(pattern)
	windowHash := h.Init(text[:m])

	for i := 0; i <= n-m; i++ {
		if windowHash == patternHash {
			metrics.HashMatches++

			match := true
			for j := 0; j < m; j++ {
				metrics.CharComparisons++
				if text[i+j] != pattern[j] {
					match = false
					break
				}
			}

			if match {
				metrics.Matches++
			} else {
				metrics.Collisions++
			}
		}

		if i < n-m {
			windowHash = h.Roll(windowHash, text[i], text[i+m])
		}
	}

	metrics.Time = time.Since(start)
	return metrics
}
