// Provide functions for CMUDict management.
package cmudict

import (
	"bufio"
	"log"
	"os"
	"path"
	"regexp"
	"strings"
)

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

// HasAccent checks if a phoneme has an accent number.
func HasAccent(phoneme string) bool {
	n := len(phoneme)
	last := phoneme[n-1]
	return last == '0' || last == '1' || last == '2'
}

// IsVowel checks if a phoneme is a vowel, including ER.
func IsVowel(phoneme string) bool {
	first := phoneme[0]
	return first == 'A' || first == 'E' || first == 'I' || first == 'O' || first == 'U'
}

// LoadDefaultDict loads the CMU dictionary file from the location specified by the CMUDICT environment variable.
func LoadDefaultDict() map[string]string {
	dir := os.Getenv("CMUDICT")
	if dir == "" {
		log.Fatal("Set CMUDICT variable to directory of dictionary file")
	}
	path := os.Join(dir, "cmudict.0.7a")
	return LoadDict(path)
}

// LoadDefaultSymbols loads the CMU symbols file from the location specified by the CMUDICT environment variable.
func LoadDefaultSymbols(accent bool) map[string]bool {
	dir := os.Getenv("CMUDICT")
	if dir == "" {
		log.Fatal("Set CMUDICT variable to directory of symbols file")
	}
	path := os.Join(dir, "cmudict.0.7a.symbols")
	return LoadSymbols(path, accent)
}

// LoadDict loads the CMU dictionary file and returns it as a map.
func LoadDict(file string) map[string]string {
	// Open file.
	handle, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Scan file line by line.
	dict := make(map[string]string)
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// Skip comments
		if line[:3] == ";;;" {
			continue
		}
		fields := strings.Split(line, "  ")
		word := fields[0]
		pron := fields[1]
		dict[word] = pron
	}
	return dict
}

// LoadSymbols loads the CMU symbols file and returns it as a map.
// This function removes the vowels symbols without accent numbers.
func LoadSymbols(file string, accent bool) map[string]bool {
	// Open file.
	handle, err := os.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()

	// Scan file line by line.
	symbols := make(map[string]bool)
	scanner := bufio.NewScanner(handle)
	for scanner.Scan() {
		phoneme := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
		if IsVowel(phoneme) {
			if HasAccent(phoneme) {
				if !accent {
					// Strip accents since requested to do so.
					phoneme = StripAccent(phoneme)
				}
			} else {
				// Remove vowel symbols that already had no accent numbers.
				continue
			}
		}
		symbols[phoneme] = true
	}
	return symbols
}

// StripAccent removes the accent number from a phoneme.
func StripAccent(phoneme string) string {
	n := len(phoneme)
	last := phoneme[n-1]
	if last == '0' || last == '1' || last == '2' {
		phoneme = phoneme[:n-1]
	}
	return phoneme
}

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
