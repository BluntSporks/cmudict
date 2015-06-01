package cmudict

import (
	"bufio"
	"log"
	"os"
)

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
