package cmudict

// StripIndex removes the index number in parentheses from the end of a word.
func StripIndex(word string) string {
	n := len(word)
	if n > 3 {
		last := word[n-1]
		if last == ')' {
			next := word[n-3]
			if next == '(' {
				word = word[:n-3]
			}
		}
	}
	return word
}
