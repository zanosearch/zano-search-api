package nlp

import (
	"github.com/kljensen/snowball"
	"slices"
)

func Stemmer(tokens []string) []string {
	var stemmedTokens []string

	for _, token := range tokens {
		stemmed, err := snowball.Stem(token, "english", true)
		if err == nil {

			// remove if statement if you just want to return stemmed without retaining original stemmed words
			if !slices.Contains(tokens, stemmed) {
				stemmedTokens = append(stemmedTokens, stemmed)
			}
		}
	}

	// keep previous unstemmed words and add stemmed
	finalTokens := append(tokens, stemmedTokens...)

	return finalTokens

}
