package main

import (
	"fmt"
	"log"
	"os"

	"hash_research/rk"
	"hash_research/hash"
	"hash_research/data"

)

func writeToFile(filename, content string) error {
	return os.WriteFile(filename, []byte(content), 0644)
}

func readFromFile(filename string) (string, error) {
	s, err := os.ReadFile(filename)
	return string(s), err
}


func main() {

	text, err := readFromFile("data/lorem.txt")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}

	var h hash.Hasher
	h = &hash.SumHash{}
	rk.Search(text, "dff", h)
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