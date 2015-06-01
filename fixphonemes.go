package cmudict

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
			} else if IsVowel(ph) && IsVowel(newPhonemes[i+1]) {
				split = true
			}
			if split {
				out = append(out, "'")
			}
		}
	}
	return out
}
