package nlp

func DefaultNlp(query string) ([]string, error) {
	tokens := Tokenizer(query)
	stopWordFiltered := StopWords(tokens)
	stemmedTokens := Stemmer(stopWordFiltered)

	return stemmedTokens, nil
}
