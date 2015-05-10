// Provide functions for CMUDict management.
package cmudict

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"

	"github.com/BluntSporks/cmudict"
)

// DefaultDictPath gets the CMU dictionary file location from the CMUDICT_DATA environment variable.
func DefaultDictPath() string {
	dir := os.Getenv("CMUDICT_DATA")
	if dir == "" {
		log.Fatal("Set CMUDICT_DATA variable to directory of dictionary file")
	}
	return path.Join(dir, "cmudict.0.7a")
}

// DefaultSymbolPath gets the CMU symbols file location from the CMUDICT_DATA environment variable.
func DefaultSymbolPath() string {
	dir := os.Getenv("CMUDICT_DATA")
	if dir == "" {
		log.Fatal("Set CMUDICT_DATA variable to directory of symbols file")
	}
	return path.Join(dir, "cmudict.0.7a.symbols")
}

// FixPhonemes fixes a pronunciation string with custom fixes designed for better phonetic spelling.
func FixPhonemes(phonemes []string) []string {
	n := len(phonemes)
	newPhonemes := make([]string, 0, 2*n)
	for i := 0; i < n; i++ {
		phoneme := phonemes[i]
		if len(phoneme) > 2 && phoneme[:2] == "ER" {
			// Use ahx r instead of erx.
			newPhonemes = append(newPhonemes, "AH"+string(phoneme[2]))
			phoneme = "R"
		} else if i < n-1 {
			if phoneme == "HH" && phonemes[i+1] == "W" {
				// Use wh instead of hw.
				phoneme = "WH"
				i++
			}
		}
		newPhonemes = append(newPhonemes, phoneme)
	}
	n = len(newPhonemes)
	out := make([]string, 0, 2*n)
	for i, ph := range newPhonemes {
		out = append(out, ph)
		if i < n-1 {
			// Use an apostrophe to split up ambiguous combinations of sounds.
			split := false
			if newPhonemes[i+1] == "HH" {
				if ph == "D" || ph == "S" || ph == "T" || ph == "W" || ph == "Z" {
					split = true
				}
			} else if ph == "N" && newPhonemes[i+1] == "G" {
				split = true
			} else if cmudict.IsVowel(ph) && cmudict.IsVowel(newPhonemes[i+1]) {
				split = true
			}
			if split {
				out = append(out, "'")
			}
		}
	}
	return out
}

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

// LoadDict loads the CMU dictionary file and returns it as a map.
func LoadDict(file string) map[string]string {
	// Open file.
	hdl, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer hdl.Close()

	// Scan file line by line.
	dict := make(map[string]string)
	scanner := bufio.NewScanner(hdl)
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		// Skip comments
		if line[:3] == ";;;" {
			continue
		}
		flds := strings.Split(line, "  ")
		word := flds[0]
		pron := flds[1]
		dict[word] = pron
	}
	return dict
}

// LoadSymbols loads the CMU symbols file and returns it as a map.
// This function removes the vowels symbols without accent numbers.
func LoadSymbols(file string, accent bool) map[string]bool {
	// Open file.
	hdl, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer hdl.Close()

	// Scan file line by line.
	symbols := make(map[string]bool)
	scanner := bufio.NewScanner(hdl)
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
