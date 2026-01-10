package main

import (
	"fmt"
	"log"
	"os"
	"encoding/csv"
	"strconv"

	"hash_research/rk"
	"hash_research/hash"
	"hash_research/data"

)

func writeToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func readFromFile(filename string) (string) {
	s, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(s)
}

func experimentTextLength(
	patternLength int,
	output string,
) {

	files := []struct {
		name string
		path string
	}{
		{"lorem", "data/lorem.txt"},
		{"dna",   "data/dna.txt"},
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
		text := readFromFile(f.path)

		start := len(text) / 2
		pattern := text[start : start+patternLength]

		for _, hf := range hashFactories {
			h := hf.new()
			m := rk.Search(text, pattern, h)

			appendResult(
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

		appendResult(
			output,
			f.name,
			bigH.Name(),
			pattern,
			text,
			mBig,
		)
	}
}


func experimentPatternLength(
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
		if start + m >= len(text) {
			continue
		}

		pattern := text[start : start+m]

		for _, hf := range hashFactories {
			h := hf.new()
			metrics := rk.Search(text, pattern, h)

			appendResult(
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

		appendResult(
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
		"lorem": readFromFile("data/lorem.txt"),
		"alice": readFromFile("data/alice.txt"),
		"dna":   readFromFile("data/dna.txt"),
		"repetitive_100k": readFromFile("data/repetitive_100k.txt"),
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


func appendResult(
	filename string,
	textName string,
	hashName string,
	pattern string,
	text string,
	m rk.Metrics) error {

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	if err := writeCSVHeaderIfNeeded(file); err != nil {
		return err
	}

	writer := csv.NewWriter(file)
	defer writer.Flush()

	return writer.Write([]string{
		textName,
		hashName,
		strconv.Itoa(len(pattern)),
		strconv.Itoa(len(text)),
		strconv.FormatInt(m.Time.Nanoseconds(), 10),
		strconv.Itoa(m.HashMatches),
		strconv.Itoa(m.Collisions),
		strconv.Itoa(m.CharComparisons),
		strconv.Itoa(m.Matches),
	})
}


func writeCSVHeaderIfNeeded(file *os.File) error {
	info, err := file.Stat()
	if err != nil {
		return err
	}

	if info.Size() == 0 {
		writer := csv.NewWriter(file)
		defer writer.Flush()

		return writer.Write([]string{
			"text_name",
			"hash_name",
			"pattern_length",
			"text_length",
			"time_ns",
			"hash_matches",
			"collisions",
			"char_comparisons",
			"matches",
		})
	}
	return nil
}

func randomTextGen() { // вызвать для генерации файлов со случайными последовательностями
	sizes := []int{
		10_000,
		50_000,
		100_000,
		200_000,
	}

	for _, size := range sizes {
		text := data.RandomText(size)
		filename := fmt.Sprintf("data/random_%d.txt", size)

		err := writeToFile(filename, text)
		if err != nil {
			panic(err)
		}

		fmt.Println("generated:", filename)
	}

	repetitive := data.RepetitiveText(100_000, 'A')
	writeToFile("data/repetitive_100k.txt", repetitive)
}