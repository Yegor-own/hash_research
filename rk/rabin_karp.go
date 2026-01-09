package rk

import (
	"time"

	"hash_research/hash"
)

type Metrics struct {
	Time            time.Duration // общее время работы
	HashMatches     int // истинные совпадения + коллизии
	Matches         int // истинные совпадения
	Collisions      int // только коллизии
	CharComparisons int // количество сравнений символов
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
		if windowHash == patternHash { // сравнение хэшей
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


func SearchBig(text, pattern string, h hash.BigHasher) Metrics {
	n, m := len(text), len(pattern)
	var metrics Metrics

	if m > n {
		return metrics
	}

	start := time.Now()

	patternHash := h.Init(pattern)
	windowHash := h.Init(text[:m])

	for i := 0; i <= n-m; i++ {

		// СРАВНЕНИЕ big.Int
		if windowHash.Cmp(patternHash) == 0 {
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