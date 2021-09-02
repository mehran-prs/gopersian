package gopersian

import (
	"strings"

	"golang.org/x/text/unicode/bidi"
)

func adjustLetterShape(previousChar, currentChar, nextChar rune) rune {
	shape := currentChar
	previousIn := false // in the Arabic Alphabet or not
	nextIn := false     // in the Arabic Alphabet or not

	for _, s := range alphabet {
		if s.equals(previousChar) { // previousChar in the Arabic Alphabet ?
			previousIn = true
		}

		if s.equals(nextChar) { // nextChar in the Persian Alphabet ?
			nextIn = true
		}
	}

	for _, s := range alphabet {
		if !s.equals(currentChar) { // currentChar in the Persian Alphabet ?
			continue
		}

		if previousIn && nextIn { // between two Persian Alphabet, return the medium shape
			for s, _ := range isolatedAfter {
				if s.equals(previousChar) {
					return letterFromRune(currentChar).Initial
				}
			}

			return letterFromRune(currentChar).Medial
		}

		if nextIn { // beginning (because the previous is not in the Arabic Alphabet)
			return letterFromRune(currentChar).Initial
		}

		if previousIn { // final (because the next is not in the Arabic Alphabet)
			for s, _ := range isolatedAfter {
				if s.equals(previousChar) {
					return letterFromRune(currentChar).Isolated
				}
			}
			return letterFromRune(currentChar).Final
		}

		if !previousIn && !nextIn {
			return letterFromRune(currentChar).Isolated
		}

	}
	return shape
}

// Shape returns the glyph representation of the given text
func Shape(val string) string {
	var prev, next rune

	runes := []rune(val)
	length := len(runes)
	newText := make([]rune, 0, length)

	for i, current := range runes {
		// get the previous char
		if (i - 1) < 0 {
			prev = 0
		} else {
			prev = runes[i-1]
		}

		// get the next char
		if (i + 1) <= length-1 {
			next = runes[i+1]
		} else {
			next = 0
		}

		// get the current char representation or return the same if unnecessary
		glyph := adjustLetterShape(prev, current, next)

		// append the new char representation to the newText
		newText = append(newText, glyph)
	}

	return string(newText)
}

// Reorder reorder runes by checking directions.
func Reorder(val string) (string, error) {
	p := new(bidi.Paragraph)
	_, err := p.SetString(val)
	if err != nil {
		return "", err
	}

	ordering, err := p.Order()
	if err != nil {
		return "", err
	}

	builder := strings.Builder{}
	olen := ordering.NumRuns()

	for i := 0; i < olen; i++ {
		index := i
		if ordering.Direction() == bidi.RightToLeft {
			index = olen - i - 1
		}
		r := ordering.Run(index)

		str := r.String()
		if r.Direction() == bidi.RightToLeft {
			str = bidi.ReverseString(r.String())
		}
		builder.WriteString(str)
	}

	return builder.String(), nil
}

// Bidi shapes and reorder the string.
func Bidi(val string) (string, error) {
	return Reorder(Shape(val))
}
