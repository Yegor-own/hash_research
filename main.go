package main

import (
	"hash_research/hash"
	"hash_research/rk"
)


func experimentTextLength( // эксперимент 2 влияние длины текста
	patternLength int,
	output string,
) {

	files := []struct {
		name string
		path string
	}{
		{"lorem", "data/lorem.txt"},
		{"dna", "data/dna.txt"},
		{"alice", "data/alice.txt"},
		{"random_10k", "data/random_10000.txt"},
		{"random_50k", "data/random_50000.txt"},
		{"random_100k", "data/random_100000.txt"},
		{"random_200k", "data/random_200000.txt"},
	}

	hashFactories := []struct {
		name string
		new  func() hash.Hasher
	}{
		{"sum", func() hash.Hasher { return &hash.SumHash{} }},
		{"poly_mod", func() hash.Hasher { return hash.NewPolyHash(257, 1_000_000_007) }},
		{"djb2", func() hash.Hasher { return &hash.DJB2Hash{} }},
	}

	for _, f := range files {
		text := ReadFromFile(f.path)

		start := len(text) / 2
		pattern := text[start : start+patternLength]

		for _, hf := range hashFactories {
			h := hf.new()
			m := rk.Search(text, pattern, h)

			AppendResult(
				output,
				f.name,
				h.Name(),
				pattern,
				text,
				m,
			)
		}

		bigH := hash.NewPolyHashBig(257)
		mBig := rk.SearchBig(text, pattern, bigH)

		AppendResult(
			output,
			f.name,
			bigH.Name(),
			pattern,
			text,
			mBig,
		)
	}
}

func experimentPatternLength( // эксперимент 1 влияние длины шаблона
	textName string,
	text string,
	output string,
) {

	patternLengths := []int{2, 4, 8, 16, 32, 64, 128, 256, 512}

	hashFactories := []struct {
		name string
		new  func() hash.Hasher
	}{
		{"sum", func() hash.Hasher { return &hash.SumHash{} }},
		{"poly_mod", func() hash.Hasher { return hash.NewPolyHash(257, 1_000_000_007) }},
		{"djb2", func() hash.Hasher { return &hash.DJB2Hash{} }},
	}

	for _, m := range patternLengths {
		start := len(text) / 2
		if start+m >= len(text) {
			continue
		}

		pattern := text[start : start+m]

		for _, hf := range hashFactories {
			h := hf.new()
			metrics := rk.Search(text, pattern, h)

			AppendResult(
				output,
				textName,
				h.Name(),
				pattern,
				text,
				metrics,
			)
		}

		// poly_nomod (big.Int)
		bigH := hash.NewPolyHashBig(257)
		metricsBig := rk.SearchBig(text, pattern, bigH)

		AppendResult(
			output,
			textName,
			bigH.Name(),
			pattern,
			text,
			metricsBig,
		)
	}
}

func main() {

	texts := map[string]string{
		"lorem":           ReadFromFile("data/lorem.txt"),
		"alice":           ReadFromFile("data/alice.txt"),
		"dna":             ReadFromFile("data/dna.txt"),
		"repetitive_100k": ReadFromFile("data/repetitive_100k.txt"),
	}

	for name, text := range texts {
		experimentPatternLength(
			name,
			text,
			"output/pattern_length.csv",
		)
	}

	experimentTextLength(32, "output/text_length.csv")
}





