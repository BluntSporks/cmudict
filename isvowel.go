package cmudict

// IsVowel checks if a phoneme is a vowel, including ER.
func IsVowel(phoneme string) bool {
	first := phoneme[0]
	return first == 'A' || first == 'E' || first == 'I' || first == 'O' || first == 'U'
}
