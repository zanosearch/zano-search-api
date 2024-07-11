package nlp

import (
	"strings"
	"unicode"
)

func Tokenizer(query string) []string {

	return strings.FieldsFunc(strings.ToLower(query), func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

}
