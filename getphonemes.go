package cmudict

import "strings"

// GetPhonemes returns the list of phonemes from a string, with or without accents.
func GetPhonemes(pron string, accent bool) []string {
	phonemes := strings.Split(pron, " ")
	if !accent {
		for i, phoneme := range phonemes {
			phonemes[i] = StripAccent(phoneme)
		}
	}
	return phonemes
}
