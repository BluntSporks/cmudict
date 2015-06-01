package cmudict

// HasAccent checks if a phoneme has an accent number.
func HasAccent(phoneme string) bool {
	n := len(phoneme)
	last := phoneme[n-1]
	return last == '0' || last == '1' || last == '2'
}
