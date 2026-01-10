package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"hash_research/data"
	"hash_research/rk"
)


func WriteToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func ReadFromFile(filename string) string {
	s, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	return string(s)
}

func AppendResult(
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

	if err := WriteCSVHeaderIfNeeded(file); err != nil {
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

func WriteCSVHeaderIfNeeded(file *os.File) error {
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

		err := WriteToFile(filename, text)
		if err != nil {
			panic(err)
		}

		fmt.Println("generated:", filename)
	}

	repetitive := data.RepetitiveText(100_000, 'A')
	WriteToFile("data/repetitive_100k.txt", repetitive)
}
