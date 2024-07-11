package nlp

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func StopWords(tokens []string) []string {
	var stopWords []string
	var newTokens []string

	file, err := os.Open("stop-words-english.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer func(file *os.File) {
		_ = file.Close()
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		stopWords = append(stopWords, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	for _, token := range tokens {
		if !slices.Contains(stopWords, token) {
			newTokens = append(newTokens, token)
		}
	}

	return newTokens
}
