package cmudict

// StripAccent removes the accent number from a phoneme.
func StripAccent(phoneme string) string {
	n := len(phoneme)
	last := phoneme[n-1]
	if last == '0' || last == '1' || last == '2' {
		phoneme = phoneme[:n-1]
	}
	return phoneme
}
