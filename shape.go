package gopersian

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

func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
